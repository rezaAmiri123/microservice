DB_NAME=microservice
DB_USER=postgres
DB_PASSWORD=postgres
DB_ADDR=microservice_postgesql
DB_URL=postgresql://$(DB_USER):$(DB_PASSWORD)@$(api_postgesql):5432/$(DB_NAME)?sslmode=disable

createdb:
	docker exec -it $(DB_ADDR) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it $(DB_ADDR) dropdb $(DB_NAME)
	
psql:
	psql -h $(DB_ADDR) -U $(DB_USER) -d $(DB_NAME) -W

evans:
	evans --host localhost --port 5080 -r repl

migratecreate:
	migrate create -ext sql -dir internal/adapters/migration -seq service_user

migrateup:
	migrate -path internal/adapters/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path internal/adapters/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path internal/adapters/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path internal/adapters/migration -database "$(DB_URL)" -verbose down 1

migrateforce:
	migrate  -path internal/adapters/migration -database "$(DB_URL)" force 1


test_coverage:
	go test --cover ./...

user_proto:
	protoc \
		--go_out=proto/grpc --go_opt=paths=source_relative \
		--go-grpc_out=proto/grpc --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=proto/grpc --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=service_user \
		--proto_path=proto/grpc proto/grpc/*.proto