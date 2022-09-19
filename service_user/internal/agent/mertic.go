package agent

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rezaAmiri123/microservice/service_user/internal/metrics"
)

func (a *Agent) setupMetric() error {
	metric := metrics.NewUserServiceMetric(&metrics.Config{
		MetricServiceName:     a.MetricServiceName,
		MetricServiceHostPort: a.MetricServiceHostPort,
	})
	//prometheus.MustRegister(metric.CreateUserHttpRequests)
	a.metric = metric
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(a.MetricServiceHostPort, nil)
	return nil
}
