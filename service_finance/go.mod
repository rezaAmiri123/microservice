module github.com/rezaAmiri123/microservice/service_finance

go 1.19

require (
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/rezaAmiri123/microservice/pkg v0.0.0-20220927124653-e36732549ab7
)

replace github.com/rezaAmiri123/microservice/pkg => ../pkg
