package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Config struct {
	MetricServiceName     string `mapstructure:"METRIC_SERVICE_NAME"`
	MetricServiceHostPort string `mapstructure:"METRIC_SERVICE_HOST_PORT"`
}

type MessageServiceMetric struct {
	CreateEmailKafkaRequests prometheus.Counter
	SuccessKafkaMessages     prometheus.Counter
	ErrorKafkaMessages       prometheus.Counter
}

func NewUserServiceMetric(cfg *Config) *MessageServiceMetric {
	return &MessageServiceMetric{
		CreateEmailKafkaRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_email_kafka_request_total", cfg.MetricServiceName),
			Help: "The total number of create email kafka requests",
		}),
		SuccessKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_kafka_processed_messages_total", cfg.MetricServiceName),
			Help: "The total number of success kafka processed messages",
		}),
		ErrorKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_kafka_processed_messages_total", cfg.MetricServiceName),
			Help: "The total number of error kafka processed messages",
		}),
	}
}
