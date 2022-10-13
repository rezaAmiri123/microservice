DB_NAME=microservice_message
DB_USER=postgres
DB_PASSWORD=postgres
DB_ADDR=microservice_postgesql
DB_URL=postgresql://$(DB_USER):$(DB_PASSWORD)@$(api_postgesql):5432/$(DB_NAME)?sslmode=disable
MIGRATION_PATH=internal/adapters/migration
createdb:
	docker exec -it $(DB_ADDR) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it $(DB_ADDR) dropdb $(DB_NAME)
	
psql:
	psql -h $(DB_ADDR) -U $(DB_USER) -d $(DB_NAME) -W

evans:
	evans --host localhost --port 5080 -r repl

# set evanse header
# header Authorization='Bearer xxxxx'

run:
	go run cmd/message/main.go

migratecreate:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq service_message

migrateup:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose down 1

migrateforce:
	migrate  -path $(MIGRATION_PATH) -database "$(DB_URL)" force 1


test_coverage:
	go test --cover ./...

mock_generate:
	go generate ./...

proto_message:
	protoc \
		--go_out=proto/kafka --go_opt=paths=source_relative \
		--go-grpc_out=proto/kafka --go-grpc_opt=paths=source_relative \
		--proto_path=proto/kafka  proto/kafka/email.proto