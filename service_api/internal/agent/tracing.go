package agent

import (
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/tracing"
)

func (a *Agent) setupTracing() error {
	if a.TracerEnable {
		tracer, closer, err := tracing.NewJaegerTracer(tracing.Config{
			TracerServiceName: a.TracerServiceName,
			TracerHostPort:    a.TracerHostPort,
			TracerEnable:      a.TracerEnable,
			TracerLogSpans:    a.TracerLogSpans,
		})
		if err != nil {
			return err
		}
		opentracing.SetGlobalTracer(tracer)
		a.closers = append(a.closers, closer)
	}
	return nil
}
