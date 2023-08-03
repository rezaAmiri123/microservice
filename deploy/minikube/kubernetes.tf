// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/namespace
resource kubernetes_namespace_v1 namespace {
  metadata {
    name = var.project
  }
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/config_map_v1
resource kubernetes_config_map_v1 common {
  metadata {
    name      = "common-config-map"
    namespace = var.project
  }

  data = {
    ENVIRONMENT  = "production"
    WEB_PORT     = ":80"
    RPC_PORT     = ":9000"
    NATS_URL     = "nats:4222"
    RPC_SERVICES = "STORES=stores:9000,CUSTOMERS=customers:9000"
  }

  depends_on = [
    kubernetes_namespace_v1.namespace
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/ingress_v1
resource kubernetes_ingress_v1 swagger {
  metadata {
    name        = "swagger-ingress"
    namespace   = var.project
    annotations = {
      "nginx.ingress.kubernetes.io/rewrite-target" = "/"
      # "alb.ingress.kubernetes.io/group.name"         = var.project
      # "alb.ingress.kubernetes.io/scheme"             = "internet-facing"
      # "alb.ingress.kubernetes.io/load-balancer-name" = var.project
      # "alb.ingress.kubernetes.io/inbound-cidrs"      = local.allowed_cidr_block
      # "alb.ingress.kubernetes.io/target-type"        = "instance"
    }
  }

  spec {
    rule {
      # host = "mallbots.local"
      http {
        path {
          path      = "/"
          path_type = "Prefix"
          backend {
            service {
              name = "users" # pick a service; any service
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
    kubernetes_service_v1.users,
  ]
 }

# output swagger_url {
#   value = "http://${kubernetes_ingress_v1.swagger.status[0].load_balancer[0].ingress[0].hostname}"
# }
