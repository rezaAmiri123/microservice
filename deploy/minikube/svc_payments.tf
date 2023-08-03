// https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/password
resource random_password payments {
  length = 16
  special = false
}

// https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource
// https://www.terraform.io/language/resources/provisioners/local-exec
resource null_resource init_payments_db {
  provisioner "local-exec" {
    command     = "psql --file sql/init_service_db.psql -v db=$DB -v user=$USER -v pass=$PASS postgres://postgres:${random_password.postgres.result}@${var.db_endpoint}:${var.db_port}/postgres"
    environment = {
      DB   = "payments"
      USER = "payments_user"
      PASS = random_password.payments.result
    }
  }
  depends_on = [
    null_resource.init_db,
    random_password.payments
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/secret_v1
resource kubernetes_secret_v1 payments {
  metadata {
    name      = "payments-secrets"
    namespace = var.project
  }

  data = {
    # PG_CONN = "host=${var.db_host} port=${var.db_port} dbname=users user=users_user password=${random_password.users.result} search_path=users,public"
    POSTGRES_PASSWORD: random_password.payments.result
  }
  depends_on = [
    kubernetes_namespace_v1.namespace,
    null_resource.init_payments_db
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/config_map_v1
resource kubernetes_config_map_v1 payments {
  metadata {
    name      = "payments-config-map"
    namespace = var.project
  }

  data = {
      POSTGRES_DRIVER= "pgx"
      POSTGRES_HOST= "postgres"
      POSTGRES_PORT= "5432"
      POSTGRES_USER= "payments_user"
      POSTGRES_DB_NAME= "payments"
      POSTGRES_SEARCH_PATH= "payments,public"
      HTTP_SERVER_PORT= "80"
      GRPC_SERVER_PORT= "9000"
      GRPC_STORE_CLIENT_ADDR= "stores"
      GRPC_STORE_CLIENT_PORT= "9000"
      EVENT_SERVER_TYPE= "nats"
      NATS_URL= "nats:4222"
      OTEL_SERVICE_NAME= "payments"
      OTEL_EXPORTER_OTLP_ENDPOINT= "http://collector:4317"
  }

  depends_on = [
    kubernetes_namespace_v1.namespace
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/deployment_v1
resource kubernetes_deployment_v1 payments {
  metadata {
    name      = "payments"
    namespace = var.project
    labels    = {
      app = "payments"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app.kubernetes.io/name" = "payments"
      }
    }
    template {
      metadata {
        name   = "payments"
        labels = {
          "app.kubernetes.io/name" = "payments"
        }
      }
      spec {
        hostname = "payments"
        container {
          name              = "payments"
          image             = "${var.image_repo}/payments:latest"
          image_pull_policy = "IfNotPresent"
          env_from {
            config_map_ref {
              name = "payments-config-map"
            }
          }
          env_from {
            secret_ref {
              name = "payments-secrets"
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
    kubernetes_config_map_v1.payments,
    kubernetes_secret_v1.payments,
    kubernetes_service_v1.nats,
    null_resource.init_db
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_v1
resource kubernetes_service_v1 payments {
  metadata {
    name      = "payments"
    namespace = var.project
    labels    = {
      app = "payments"
    }
  }
  spec {
    selector = {
      "app.kubernetes.io/name" = "payments"
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
resource kubernetes_ingress_v1 payments {
  metadata {
    name        = "payments-ingress"
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
          path      = "/v1/api/payments/"
          path_type = "Prefix"
          backend {
            service {
              name = "payments"
              port {
                number = 80
              }
            }
          }
        }
        path {
          path      = "/payments-spec/"
          path_type = "Prefix"
          backend {
            service {
              name = "payments"
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
    kubernetes_service_v1.payments,
  ]
}
