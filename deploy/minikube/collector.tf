// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/config_map_v1
resource kubernetes_config_map_v1 collector {
  metadata {
    name      = "collector-config-map"
    namespace = var.project
  }

  data = {
    "otel-config.yml"  = <<EOF
    receivers:
      otlp:
        protocols:
          grpc:
      otlp/spanmetrics:
        protocols:
          grpc:
            endpoint: "localhost:65535"

    exporters:
      jaeger:
        endpoint: "jaeger:14250"
        tls:
          insecure: true
      logging:
      prometheus:
        endpoint: "collector:9464"


    # endpoint: "prometheus:9090"
    processors:
      batch:
      spanmetrics:
        metrics_exporter: prometheus

    service:
      pipelines:
        traces:
          receivers: [ otlp ]
          processors: [ spanmetrics, batch ]
          exporters: [ logging, jaeger ]
        metrics:
          receivers: [ otlp ]
          processors: [ batch ]
          exporters: [ prometheus, logging ]
        metrics/spanmetrics:
          receivers: [ otlp/spanmetrics ]
          exporters: [ prometheus, logging ]
    EOF
  }

  depends_on = [
    kubernetes_namespace_v1.namespace
  ]
}


// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/deployment_v1
resource kubernetes_deployment_v1 collector {
  metadata {
    name      = "collector"
    namespace = var.project
    labels    = {
      app = "collector"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app.kubernetes.io/name" = "collector"
      }
    }
    template {
      metadata {
        name   = "collector"
        labels = {
          "app.kubernetes.io/name" = "collector"
        }
      }
      spec {
        hostname       = "collector"
        restart_policy = "Always"
        container {
          image = "otel/opentelemetry-collector-contrib:0.60.0"
          image_pull_policy = "IfNotPresent"
          name  = "collector"
          args  = ["--config","/tmp/otel/otel-config.yml"]
          port {
            # pprof extension
            protocol       = "TCP"
            container_port = 1888
          }
          port {
            # Prometheus metrics exposed by the collector
            protocol       = "TCP"
            container_port = 8888
          }
          port {
            # Prometheus exporter metrics
            protocol       = "TCP"
            container_port = 8889
          }
          port {
            # health_check extension
            protocol       = "TCP"
            container_port = 13133
          }
          port {
            # OTLP gRPC receiver
            protocol       = "TCP"
            container_port = 4317
          }
          port {
            # OTLP http receiver
            protocol       = "TCP"
            container_port = 4318
          }
          volume_mount {
            mount_path = "/tmp/otel/"
            name       = "oteldata"
          }

        }
        volume {
          name = "oteldata"
          config_map {
            name = "collector-config-map"
            default_mode = "0777"
          }
        }
      }
    }
  }
  depends_on = [
    kubernetes_namespace_v1.namespace,
    kubernetes_config_map_v1.collector,
    kubernetes_service_v1.jaeger
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_v1
resource kubernetes_service_v1 collector {
  metadata {
    name      = "collector"
    namespace = var.project
    labels    = {
      app = "collector"
    }
  }
  spec {
    selector = {
      "app.kubernetes.io/name" = "collector"
    }
    port {
      name = "port-1"
      protocol    = "TCP"
      port        = 1888
      target_port = 1888
    }
    port {
      name = "port-2"
      protocol    = "TCP"
      port        = 8888
      target_port = 8888
    }
    port {
      name = "port-3"
      protocol    = "TCP"
      port        = 8889
      target_port = 8889
    }
    port {
      name = "port-4"
      protocol    = "TCP"
      port        = 13133
      target_port = 13133
    }
    port {
      name = "port-5"
      protocol    = "TCP"
      port        = 4317
      target_port = 4317
    }
    port {
      name = "port-6"
      protocol    = "TCP"
      port        = 4318
      target_port = 4318
    }
    type = "ClusterIP"
  }
  depends_on = [
    kubernetes_namespace_v1.namespace,
    kubernetes_deployment_v1.collector
  ]
}
