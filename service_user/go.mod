module github.com/rezaAmiri123/microservice/service_user

go 1.18

require (
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0
	github.com/rezaAmiri123/microservice/pkg v0.0.0-20220915112118-0dbdbf6dfa4d
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
)

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3
	github.com/jmoiron/sqlx v1.3.5
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/client_golang v1.13.0
	github.com/stretchr/testify v1.8.0
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sys v0.0.0-20220610221304-9f5ed59c137d // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//go mod edit -replace github.com/rezaAmiri123/microservice/pkg ../pkg

replace github.com/rezaAmiri123/microservice/pkg => ../pkg
