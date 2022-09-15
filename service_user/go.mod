module github.com/rezaAmiri123/microservice/service_user

go 1.18

require (
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0
	github.com/rezaAmiri123/microservice/pkg v0.0.0-20220915112118-0dbdbf6dfa4d
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
)

require (
	github.com/opentracing/opentracing-go v1.2.0
	github.com/stretchr/testify v1.8.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//go mod edit -replace github.com/rezaAmiri123/microservice/pkg ../pkg

replace github.com/rezaAmiri123/microservice/pkg => ../pkg
