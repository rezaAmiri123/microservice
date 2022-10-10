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

type UserServiceMetric struct {
	CreateUserGrpcRequests prometheus.Counter
	UpdateUserGrpcRequests prometheus.Counter
	LoginRequests          prometheus.Counter
	LoginVerifyRequests    prometheus.Counter
	SuccessGrpcRequests    prometheus.Counter
	ErrorGrpcRequests      prometheus.Counter
	// SuccessKafkaMessages prometheus.Counter
	// ErrorKafkaMessages   prometheus.Counter
}

func NewUserServiceMetric(cfg *Config) *UserServiceMetric {
	return &UserServiceMetric{
		CreateUserGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_user_grpc_requests_total", cfg.MetricServiceName),
			Help: "The total of create user grpc requests",
		}),
		UpdateUserGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_user_grpc_requests_total", cfg.MetricServiceName),
			Help: "The total of update user grpc requests",
		}),
		LoginRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_login_grpc_requests_total", cfg.MetricServiceName),
			Help: "The total of login grpc requests",
		}),
		LoginVerifyRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_login_verify_grpc_requests_total", cfg.MetricServiceName),
			Help: "The total of login verify grpc requests",
		}),
		SuccessGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_grpc_requsts_total", cfg.MetricServiceName),
			Help: "The total number of success grpc requests",
		}),
		ErrorGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_grpc_requsts_total", cfg.MetricServiceName),
			Help: "The total number of error grpc requests",
		}),
		// SuccessKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
		// 	Name: fmt.Sprintf("%s_success_kafka_processed_messages_total", cfg.ServiceName),
		// 	Help: "The total number of success kafka processed messages",
		// }),
		// ErrorKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
		// 	Name: fmt.Sprintf("%s_error_kafka_processed_messages_total", cfg.ServiceName),
		// 	Help: "The total number of error kafka processed messages",
		// }),
	}
}
