module github.com/rezaAmiri123/microservice/service_api

go 1.18

replace github.com/rezaAmiri123/microservice/pkg => ../pkg

replace github.com/rezaAmiri123/microservice/service_user => ../service_user

replace github.com/rezaAmiri123/microservice/service_finance => ../service_finance

require (
	github.com/google/uuid v1.3.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/rezaAmiri123/microservice/service_user v0.0.0-00010101000000-000000000000
	github.com/rezaAmiri123/test-microservice v0.0.0-20220908112505-ed25ec411f4a
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/segmentio/kafka-go v0.4.35 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/net v0.0.0-20220909164309-bea034e7d591 // indirect
	golang.org/x/sys v0.0.0-20220926163933-8cfa568d3c25 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220902135211-223410557253 // indirect
	google.golang.org/grpc v1.49.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
