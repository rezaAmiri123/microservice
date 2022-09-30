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

type FinanceServiceMetric struct {
	CreateAccountGrpcRequests  prometheus.Counter
	CreateTransferGrpcRequests prometheus.Counter
	// UpdateUserGrpcRequests prometheus.Counter
	// LoginRequests          prometheus.Counter
	SuccessGrpcRequests prometheus.Counter
	ErrorGrpcRequests   prometheus.Counter
}

func NewFinanceServiceMetric(cfg *Config) *FinanceServiceMetric {
	return &FinanceServiceMetric{
		CreateAccountGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_account_grpc_requests_total", cfg.MetricServiceHostPort),
			Help: "The total of create account grpc requests",
		}),
		CreateTransferGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_transfer_grpc_requests_total", cfg.MetricServiceHostPort),
			Help: "The total of create transfer grpc requests",
		}),
		// UpdateUserGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
		// 	Name: fmt.Sprintf("%s_update_user_grpc_requests_total", cfg.MetricServiceHostPort),
		// 	Help: "The total of update user grpc requests",
		// }),
		// LoginRequests: promauto.NewCounter(prometheus.CounterOpts{
		// 	Name: fmt.Sprintf("%s_login_grpc_requests_total", cfg.MetricServiceHostPort),
		// 	Help: "The total of login grpc requests",
		// }),
		SuccessGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_grpc_requsts_total", cfg.MetricServiceHostPort),
			Help: "The total number of success grpc requests",
		}),
		ErrorGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_grpc_requsts_total", cfg.MetricServiceHostPort),
			Help: "The total number of error grpc requests",
		}),
	}
}
