NAME=notifications
DB_NAME=$(NAME)
DB_USER=notifications_user
DB_PASSWORD=notifications_pass
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
	echo done

proto:
	protoc \
		--go_out=notificationspb --go_opt=paths=source_relative \
		--go-grpc_out=notificationspb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=notificationspb --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=internal/ports/rest/ --openapiv2_opt=allow_merge=true,merge_file_name=api \
		--proto_path=notificationspb notificationspb/*.proto

createdb:
	docker exec -it $(DB_ADDR) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

migratecreate:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(NAME)

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
mongo:
	mongosh -u admin -p admin
# db.users_cache.find()

generate:
	mockery --quiet --dir ./internal -r --all --inpackage --case underscore
	mockery --quiet --dir ./notificationspb -r --all --inpackage --case underscore
