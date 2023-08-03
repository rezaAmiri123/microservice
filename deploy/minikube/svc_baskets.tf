// https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/password
resource random_password baskets {
  length = 16
  special = false
}

// https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource
// https://www.terraform.io/language/resources/provisioners/local-exec
resource null_resource init_baskets_db {
  provisioner "local-exec" {
    command     = "psql --file sql/init_service_db.psql -v db=$DB -v user=$USER -v pass=$PASS postgres://postgres:${random_password.postgres.result}@${var.db_endpoint}:${var.db_port}/postgres"
    environment = {
      DB   = "baskets"
      USER = "baskets_user"
      PASS = random_password.baskets.result
    }
  }
  depends_on = [
    null_resource.init_db,
    random_password.baskets
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/secret_v1
resource kubernetes_secret_v1 baskets {
  metadata {
    name      = "baskets-secrets"
    namespace = var.project
  }

  data = {
    # PG_CONN = "host=${var.db_host} port=${var.db_port} dbname=users user=users_user password=${random_password.users.result} search_path=users,public"
    POSTGRES_PASSWORD: random_password.baskets.result
  }
  depends_on = [
    kubernetes_namespace_v1.namespace,
    null_resource.init_baskets_db
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/config_map_v1
resource kubernetes_config_map_v1 baskets {
  metadata {
    name      = "baskets-config-map"
    namespace = var.project
  }

  data = {
      POSTGRES_DRIVER= "pgx"
      POSTGRES_HOST= "postgres"
      POSTGRES_PORT= "5432"
      POSTGRES_USER= "baskets_user"
      POSTGRES_DB_NAME= "baskets"
      POSTGRES_SEARCH_PATH= "baskets,public"
      HTTP_SERVER_PORT= "80"
      GRPC_SERVER_PORT= "9000"
      GRPC_STORE_CLIENT_ADDR= "stores"
      GRPC_STORE_CLIENT_PORT= "9000"
      EVENT_SERVER_TYPE= "nats"
      NATS_URL= "nats:4222"
      OTEL_SERVICE_NAME= "baskets"
      OTEL_EXPORTER_OTLP_ENDPOINT= "http://collector:4317"
  }

  depends_on = [
    kubernetes_namespace_v1.namespace
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/deployment_v1
resource kubernetes_deployment_v1 baskets {
  metadata {
    name      = "baskets"
    namespace = var.project
    labels    = {
      app = "baskets"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app.kubernetes.io/name" = "baskets"
      }
    }
    template {
      metadata {
        name   = "baskets"
        labels = {
          "app.kubernetes.io/name" = "baskets"
        }
      }
      spec {
        hostname = "baskets"
        container {
          name              = "baskets"
          image             = "${var.image_repo}/baskets:latest"
          image_pull_policy = "IfNotPresent"
          env_from {
            config_map_ref {
              name = "baskets-config-map"
            }
          }
          env_from {
            secret_ref {
              name = "baskets-secrets"
            }
          }
          port {
            protocol       = "TCP"
            container_port = 80
          }
          port {
            protocol       = "TCP"
            container_port = 9000
          }
          liveness_probe {
            http_get {
              path = "/liveness"
              port = 80
            }
            initial_delay_seconds = 3
            period_seconds        = 5
          }
        }
      }
    }
  }

  depends_on = [
    kubernetes_namespace_v1.namespace,
    kubernetes_config_map_v1.baskets,
    kubernetes_secret_v1.baskets,
    kubernetes_service_v1.nats,
    null_resource.init_db
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_v1
resource kubernetes_service_v1 baskets {
  metadata {
    name      = "baskets"
    namespace = var.project
    labels    = {
      app = "baskets"
    }
  }
  spec {
    selector = {
      "app.kubernetes.io/name" = "baskets"
    }
    session_affinity = "ClientIP"
    port {
      name        = "http"
      protocol    = "TCP"
      port        = 80
      target_port = 80
    }
    port {
      name        = "grpc"
      protocol    = "TCP"
      port        = 9000
      target_port = 9000
    }
    type = "ClusterIP"
  }
  depends_on = [
    kubernetes_namespace_v1.namespace,
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/ingress_v1
resource kubernetes_ingress_v1 baskets {
  metadata {
    name        = "baskets-ingress"
    namespace   = var.project
    annotations = {
      # "alb.ingress.kubernetes.io/group.name"    = var.project
      # "alb.ingress.kubernetes.io/scheme"        = "internet-facing"
      # "alb.ingress.kubernetes.io/inbound-cidrs" = var.allowed_cidr_block
      # "alb.ingress.kubernetes.io/target-type"   = "instance"
    }
  }

  spec {
    rule {
      # host = "mallbots.local"
      http {
        path {
          path      = "/v1/api/baskets/"
          path_type = "Prefix"
          backend {
            service {
              name = "baskets"
              port {
                number = 80
              }
            }
          }
        }
        path {
          path      = "/baskets-spec/"
          path_type = "Prefix"
          backend {
            service {
              name = "baskets"
              port {
                number = 80
              }
            }
          }
        }
      }
    }
    # ingress_class_name = "alb"
  }
  depends_on = [
    kubernetes_service_v1.baskets,
  ]
}
