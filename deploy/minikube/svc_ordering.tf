// https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/password
resource random_password ordering {
  length = 16
  special = false
}

// https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource
// https://www.terraform.io/language/resources/provisioners/local-exec
resource null_resource init_ordering_db {
  provisioner "local-exec" {
    command     = "psql --file sql/init_service_db.psql -v db=$DB -v user=$USER -v pass=$PASS postgres://postgres:${random_password.postgres.result}@${var.db_endpoint}:${var.db_port}/postgres"
    environment = {
      DB   = "ordering"
      USER = "ordering_user"
      PASS = random_password.ordering.result
    }
  }
  depends_on = [
    null_resource.init_db,
    random_password.ordering
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/secret_v1
resource kubernetes_secret_v1 ordering {
  metadata {
    name      = "ordering-secrets"
    namespace = var.project
  }

  data = {
    # PG_CONN = "host=${var.db_host} port=${var.db_port} dbname=users user=users_user password=${random_password.users.result} search_path=users,public"
    POSTGRES_PASSWORD: random_password.ordering.result
  }
  depends_on = [
    kubernetes_namespace_v1.namespace,
    null_resource.init_ordering_db
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/config_map_v1
resource kubernetes_config_map_v1 ordering {
  metadata {
    name      = "ordering-config-map"
    namespace = var.project
  }

  data = {
      POSTGRES_DRIVER= "pgx"
      POSTGRES_HOST= "postgres"
      POSTGRES_PORT= "5432"
      POSTGRES_USER= "ordering_user"
      POSTGRES_DB_NAME= "ordering"
      POSTGRES_SEARCH_PATH= "ordering,public"
      HTTP_SERVER_PORT= "80"
      GRPC_SERVER_PORT= "9000"
      GRPC_STORE_CLIENT_ADDR= "stores"
      GRPC_STORE_CLIENT_PORT= "9000"
      EVENT_SERVER_TYPE= "nats"
      NATS_URL= "nats:4222"
      OTEL_SERVICE_NAME= "ordering"
      OTEL_EXPORTER_OTLP_ENDPOINT= "http://collector:4317"
  }

  depends_on = [
    kubernetes_namespace_v1.namespace
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/deployment_v1
resource kubernetes_deployment_v1 ordering {
  metadata {
    name      = "ordering"
    namespace = var.project
    labels    = {
      app = "ordering"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app.kubernetes.io/name" = "ordering"
      }
    }
    template {
      metadata {
        name   = "ordering"
        labels = {
          "app.kubernetes.io/name" = "ordering"
        }
      }
      spec {
        hostname = "ordering"
        container {
          name              = "ordering"
          image             = "${var.image_repo}/ordering:latest"
          image_pull_policy = "IfNotPresent"
          env_from {
            config_map_ref {
              name = "ordering-config-map"
            }
          }
          env_from {
            secret_ref {
              name = "ordering-secrets"
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
    kubernetes_config_map_v1.ordering,
    kubernetes_secret_v1.ordering,
    kubernetes_service_v1.nats,
    null_resource.init_db
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_v1
resource kubernetes_service_v1 ordering {
  metadata {
    name      = "ordering"
    namespace = var.project
    labels    = {
      app = "ordering"
    }
  }
  spec {
    selector = {
      "app.kubernetes.io/name" = "ordering"
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
resource kubernetes_ingress_v1 ordering {
  metadata {
    name        = "ordering-ingress"
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
          path      = "/v1/api/ordering/"
          path_type = "Prefix"
          backend {
            service {
              name = "ordering"
              port {
                number = 80
              }
            }
          }
        }
        path {
          path      = "/ordering-spec/"
          path_type = "Prefix"
          backend {
            service {
              name = "ordering"
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
    kubernetes_service_v1.ordering,
  ]
}
