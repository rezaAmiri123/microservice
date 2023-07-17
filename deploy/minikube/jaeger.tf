// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/deployment_v1
resource kubernetes_deployment_v1 jaeger {
  metadata {
    name      = "jaeger"
    namespace = var.project
    labels    = {
      app = "jaeger"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app.kubernetes.io/name" = "jaeger"
      }
    }
    template {
      metadata {
        name   = "jaeger"
        labels = {
          "app.kubernetes.io/name" = "jaeger"
        }
      }
      spec {
        hostname       = "jaeger"
        restart_policy = "Always"
        container {
          image = "jaegertracing/all-in-one:1"
          name  = "jaeger"
          image_pull_policy = "IfNotPresent"
          port {
            protocol       = "TCP"
            container_port = 16686
          }
          port {
            protocol       = "TCP"
            container_port = 14250
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
resource kubernetes_service_v1 jaeger {
  metadata {
    name      = "jaeger"
    namespace = var.project
    labels    = {
      app = "jaeger"
    }
  }
  spec {
    selector = {
      "app.kubernetes.io/name" = "jaeger"
    }
    port {
      name = "port-1"
      protocol    = "TCP"
      port        = 16686
      target_port = 16686
      node_port    = 30001
    }
    port {
      name = "port-2"
      protocol    = "TCP"
      port        = 14250
      target_port = 14250
    }
    type = "NodePort"
  }
  depends_on = [
    kubernetes_namespace_v1.namespace,
    kubernetes_deployment_v1.jaeger
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/ingress_v1
# resource kubernetes_ingress_v1 jaeger {
#   metadata {
#     name        = "jaeger-ingress"
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
#           path      = "/jaeger/"
#           path_type = "Prefix"
#           backend {
#             service {
#               name = "jaeger"
#               port {
#                 number = 16686
#               }
#             }
#           }
#         }
#       }
#     }
#     # ingress_class_name = "alb"
#   }
#   depends_on = [
#     kubernetes_service_v1.jaeger
#   ]
# }
