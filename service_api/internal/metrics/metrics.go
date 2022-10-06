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

type ApiServiceMetric struct {
	CreateUserHttpRequests    prometheus.Counter
	CreateAccountHttpRequests prometheus.Counter
	SuccessHttpRequests       prometheus.Counter
	ErrorHttpRequests         prometheus.Counter
}

func NewApiServiceMetric(cfg *Config) *ApiServiceMetric {
	return &ApiServiceMetric{
		CreateUserHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_user_http_requests_total", cfg.MetricServiceHostPort),
			Help: "The total of create user http requests",
		}),
		CreateAccountHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_account_http_requests_total", cfg.MetricServiceHostPort),
			Help: "The total of create account http requests",
		}),
		SuccessHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_http_requsts_total", cfg.MetricServiceHostPort),
			Help: "The total number of success http requests",
		}),
		ErrorHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_http_requsts_total", cfg.MetricServiceHostPort),
			Help: "The total number of error http requests",
		}),
	}
}
