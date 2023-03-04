DB_NAME=users
DB_USER=postgres
DB_PASSWORD=postgres
DB_ADDR=microservice_postgesql
DB_URL=postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_ADDR):5432/$(DB_NAME)?sslmode=disable
MIGRATION_PATH=internal/adapters/migrations


install-tools:
	echo installing tools
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/vektra/mockery/v2@latest
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	go install github.com/cucumber/godog/cmd/godog@latest
	echo done

generate:
	echo running code generation
	go generate ./...
	echo done

evans:
	evans --host localhost --port 5080 -r repl
	# evans --host 192.168.39.213 --port 30000 -r repl

evans_tls:
	evans --host localhost --port 5080 -t --servername localhost --cacert tls/ca.pem --cert tls/client.pem --certkey tls/client-key.pem -r repl

# set evanse header
# header Authorization='Bearer xxxxx'

user_proto:
	protoc \
		--go_out=userspb --go_opt=paths=source_relative \
		--go-grpc_out=userspb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=userspb --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=internal/ports/rest/ --openapiv2_opt=allow_merge=true,merge_file_name=api \
		--proto_path=userspb userspb/*.proto

migratecreate:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq service_user

migrateup:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose up
	# postgresql://postgres:postgres@192.168.39.213:30001/microservice?sslmode=disable

migrateup1:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose down 1

migrateforce:
	migrate  -path $(MIGRATION_PATH) -database "$(DB_URL)" force 1

psql:
	psql -h $(DB_ADDR) -U $(DB_USER) -d $(DB_NAME) -W
