DB_NAME=microservice_finance
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
	evans --host localhost --port 6080 -r repl
	# evans --host 192.168.39.213 --port 30002 -r repl

# set evanse header
# header Authorization='Bearer xxxxx'
# header authorization_payload='[{"id":"f5d7aec3-e573-4606-b429-8644cb264b98","username":"reza21111128","user_id":"23cfa9ca-46c5-408e-a4ee-97797b312a7b","issued_at":"2022-11-01T12:15:13.940886953Z","expired_at":"2022-11-01T12:30:13.940887121Z"}]'

run:
	go run cmd/finance/main.go

migratecreate:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq service_finance

migrateup:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose up
	# postgresql://postgres:postgres@192.168.39.213:30001/microservice_finance?sslmode=disable

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

finance_proto:
	protoc \
		--go_out=proto/grpc --go_opt=paths=source_relative \
		--go-grpc_out=proto/grpc --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=proto/grpc --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=service_finance \
		--proto_path=proto/grpc proto/grpc/*.proto

mock_generate:
	go generate ./...

make_image:
	docker build -f Dockerfile -t reza879/service_finance:$(git rev-parse --short HEAD) .
	docker push reza879/service_finance:$(git rev-parse --short HEAD)