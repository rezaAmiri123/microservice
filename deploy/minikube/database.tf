// https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/password
resource random_password postgres {
  length = 16
  special = false
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/secret_v1
resource kubernetes_secret_v1 postgres {
  metadata {
    name      = "postgres-secrets"
    namespace = var.project
  }

  data = {
    POSTGRES_PASSWORD = "${random_password.postgres.result}"
  }
  depends_on = [
    kubernetes_namespace_v1.namespace,
    random_password.postgres
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/stateful_set_v1
resource kubernetes_stateful_set_v1 postgres {
  metadata {
    name      = "postgres"
    namespace = var.project
    labels    = {
      app = "postgres"
    }
  }
  
  spec {
    pod_management_policy  = "Parallel"
    replicas               = 1
    revision_history_limit = 5

    selector {
      match_labels = {
        "app.kubernetes.io/name" = "postgres"
      }
    }
    service_name = "postgres"
    template {
      metadata {
        name   = "postgres"
        labels = {
          "app.kubernetes.io/name" = "postgres"
        }
      }
      spec {
        hostname       = "postgres"
        restart_policy = "Always"
        container {
          name              = "postgres"
          image             = "postgres:12-alpine"
          image_pull_policy = "IfNotPresent"
          env_from {
            secret_ref {
              name = "postgres-secrets"
            }
          }

          port {
            protocol       = "TCP"
            container_port = 5432
          }
          volume_mount {
            mount_path = "/var/lib/postgresql/data"
            name       = "postgresqldata"
          }
           
        
      }
      volume {
          name = "postgresqldata"
          persistent_volume_claim {
            claim_name = "postgresqldata"
          }
      }
    }
  }
  }
  
  depends_on = [
    kubernetes_namespace_v1.namespace,
  ]
}
 
// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_v1
resource kubernetes_service_v1 postgres {
  metadata {
    name      = "postgres"
    namespace = var.project
    labels    = {
      app = "postgres"
    }
  }
  spec {
    selector = {
      "app.kubernetes.io/name" = "postgres"
    }
    port {
      protocol    = "TCP"
      port        = 5432
      target_port = 5432
      node_port    = 30000
    }
    type = "NodePort"
  }
  depends_on = [
    kubernetes_namespace_v1.namespace,
    kubernetes_stateful_set_v1.postgres
  ]
}


// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/persistent_volume_claim_v1
resource kubernetes_persistent_volume_claim_v1 db {
  metadata {
    name      = "postgresqldata"
    namespace = var.project
  }
  spec {
    access_modes = ["ReadWriteOnce"]
    resources {
      requests = {
        storage : "500Mi"
      }
    }
  }
  depends_on = [
    kubernetes_namespace_v1.namespace
  ]
}

resource null_resource init_db {
  provisioner "local-exec" {
    command = "sleep 5; psql --file sql/init_db.psql postgres://postgres:${random_password.postgres.result}@${var.db_endpoint}:${var.db_port}/postgres"
  }
  depends_on = [
    kubernetes_service_v1.postgres
  ]
}


output db_password {
  value     = random_password.postgres
  sensitive = true
}

output db_conn {
  value     = "postgres://postgres:${random_password.postgres.result}@${var.db_endpoint}:${var.db_port}"
  sensitive = true
}
