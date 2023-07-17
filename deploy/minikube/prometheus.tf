// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/config_map_v1
resource kubernetes_config_map_v1 prometheus {
  metadata {
    name      = "prometheus-config-map"
    namespace = var.project
  }

  data = {
    "prometheus-config.yml"  = <<EOF
      global:
        evaluation_interval: 30s
        scrape_interval: 5s
      scrape_configs:
        - job_name: baskets
          scrape_interval: 10s
          static_configs:
            - targets:
              - 'baskets:80'
        - job_name: cosec
          scrape_interval: 10s
          static_configs:
            - targets:
              - 'cosec:80'
        - job_name: users
          scrape_interval: 10s
          static_configs:
            - targets:
              - 'users:80'
        - job_name: depot
          scrape_interval: 10s
          static_configs:
            - targets:
              - 'depot:80'
        - job_name: notifications
          scrape_interval: 10s
          static_configs:
            - targets:
              - 'notifications:80'
        - job_name: ordering
          scrape_interval: 10s
          static_configs:
            - targets:
              - 'ordering:80'
        - job_name: payments
          scrape_interval: 10s
          static_configs:
            - targets:
              - 'payments:80'
        - job_name: search
          scrape_interval: 10s
          static_configs:
            - targets:
              - 'search:80'
        - job_name: stores
          scrape_interval: 10s
          static_configs:
            - targets:
              - 'stores:80'
        - job_name: otel
          static_configs:
            - targets:
                - 'collector:9464'
        - job_name: otel-collector
          static_configs:
            - targets:
                - 'collector:8888'
    EOF
  }

  depends_on = [
    kubernetes_namespace_v1.namespace
  ]
}


// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/deployment_v1
resource kubernetes_deployment_v1 prometheus {
  metadata {
    name      = "prometheus"
    namespace = var.project
    labels    = {
      app = "prometheus"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app.kubernetes.io/name" = "prometheus"
      }
    }
    template {
      metadata {
        name   = "prometheus"
        labels = {
          "app.kubernetes.io/name" = "prometheus"
        }
      }
      spec {
        hostname       = "prometheus"
        restart_policy = "Always"
        container {
          image = "prom/prometheus:v2.37.1"
          image_pull_policy = "IfNotPresent"
          name  = "prometheus"
          args  = ["--config.file","/etc/prometheus/prometheus-config.yml"]
          port {
            protocol       = "TCP"
            container_port = 9090
          }
          volume_mount {
            mount_path = "/etc/prometheus/"
            name       = "prometheusdata"
          }

        }
        volume {
          name = "prometheusdata"
          config_map {
            name = "prometheus-config-map"
            # default_mode = "0777"
          }
        }
      }
    }
  }
  depends_on = [
    kubernetes_namespace_v1.namespace,
    kubernetes_config_map_v1.prometheus
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_v1
resource kubernetes_service_v1 prometheus {
  metadata {
    name      = "prometheus"
    namespace = var.project
    labels    = {
      app = "prometheus"
    }
  }
  spec {
    selector = {
      "app.kubernetes.io/name" = "prometheus"
    }
    port {
      protocol    = "TCP"
      port        = 9090
      target_port = 9090
      node_port    = 30002
    }
    type = "NodePort"
  }
  depends_on = [
    kubernetes_namespace_v1.namespace,
    kubernetes_deployment_v1.prometheus
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/ingress_v1
# resource kubernetes_ingress_v1 prometheus {
#   metadata {
#     name        = "prometheus-ingress"
#     namespace   = var.project
#     annotations = {
#       # "alb.ingress.kubernetes.io/group.name"    = var.project
#       # "alb.ingress.kubernetes.io/scheme"        = "internet-facing"
#       # "alb.ingress.kubernetes.io/inbound-cidrs" = var.allowed_cidr_block
#       # "alb.ingress.kubernetes.io/target-type"   = "instance"
#     }
#   }

#   spec {
#     rule {
#       # host = "mallbots.local"
#       http {
#         path {
#           path      = "/prometheus/"
#           path_type = "Prefix"
#           backend {
#             service {
#               name = "prometheus"
#               port {
#                 number = 9090
#               }
#             }
#           }
#         }
#       }
#     }
#     # ingress_class_name = "alb"
#   }
#   depends_on = [
#     kubernetes_service_v1.prometheus
#   ]
# }
