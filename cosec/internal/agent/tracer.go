package agent

import (
	"context"
	"github.com/rezaAmiri123/microservice/cosec/internal/constants"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func (a *Agent) setupTracer() error {
	exporter, err := otlptracegrpc.New(context.Background(), otlptracegrpc.WithInsecure())
	if err != nil {
		return err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		//sdktrace.WithSampler(sdktrace.AlwaysSample()),
		//sdktrace.WithResource(resource.NewSchemaless(attribute.String("service.name", "myService"))),
		//sdktrace.WithSyncer(exp),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	a.container.AddSingleton(constants.TracerKey, func(c di.Container) (any, error) {
		return tp, nil
	})

	return nil
}

//func (a *Agent) setupTracer() error {
//	l := log.New(os.Stdout, "", 0)
//	// Write telemetry data to a file.
//	f, err := os.Create("traces.txt")
//	if err != nil {
//		l.Fatal(err)
//	}
//	a.closers = append(a.closers, CloserFunc(func() error {
//		f.Close()
//		return nil
//	}))
//
//	exp, err := newExporter(f)
//	if err != nil {
//		l.Fatal(err)
//	}
//
//	exporter, err := otlptracegrpc.New(context.Background())
//	if err != nil {
//		return err
//	}
//
//	tp := sdktrace.NewTracerProvider(
//		sdktrace.WithBatcher(exporter),
//		//sdktrace.WithBatcher(exp),
//		sdktrace.WithSampler(sdktrace.AlwaysSample()),
//		sdktrace.WithResource(newResource()),
//		//sdktrace.WithSyncer(exp),
//	)
//
//	otel.SetTracerProvider(tp)
//	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
//
//	trce := tp.Tracer("xxxx")
//	_, span := trce.Start(context.Background(), "dddddddd")
//	span.End()
//	a.container.AddSingleton(constants.TracerKey, func(c di.Container) (any, error) {
//		return tp, nil
//	})
//
//	return nil
//}

// newExporter returns a console exporter.
//func newExporter(w io.Writer) (sdktrace.SpanExporter, error) {
//	return stdouttrace.New(
//		stdouttrace.WithWriter(w),
//		// Use human-readable output.
//		stdouttrace.WithPrettyPrint(),
//		// Do not print timestamps for the demo.
//		stdouttrace.WithoutTimestamps(),
//	)
//}

// newResource returns a resource describing this application.
//func newResource() *resource.Resource {
//	r, _ := resource.Merge(
//		resource.Default(),
//		resource.NewWithAttributes(
//			semconv.SchemaURL,
//			semconv.ServiceName("fib"),
//			semconv.ServiceVersion("v0.1.0"),
//			attribute.String("environment", "demo"),
//		),
//	)
//	return r
//}
