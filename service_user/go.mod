module github.com/rezaAmiri123/microservice/service_user

go 1.18

require (
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0
	github.com/rezaAmiri123/microservice/pkg v0.0.0-00010101000000-000000000000
)

//go mod edit -replace github.com/rezaAmiri123/microservice/pkg ../pkg

replace github.com/rezaAmiri123/microservice/pkg => ../pkg
