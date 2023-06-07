DB_NAME=payments
DB_USER=payments_user
DB_PASSWORD=payments_pass
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
		--go_out=paymentspb --go_opt=paths=source_relative \
		--go-grpc_out=paymentspb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=paymentspb --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=internal/ports/rest/ --openapiv2_opt=allow_merge=true,merge_file_name=api \
		--proto_path=paymentspb paymentspb/*.proto

createdb:
	docker exec -it $(DB_ADDR) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

migratecreate:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq migrate

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

generate:
	mockery --quiet --dir ./internal -r --all --inpackage --case underscore
	mockery --quiet --dir ./paymentspb -r --all --inpackage --case underscore
	swagger generate client -q -f ./internal/ports/rest/api.swagger.json -c paymentsclient -m paymentsclient/models --with-flatten=remove-unused

test_coverage:
	go test --cover ./...

test_coverage_html:
	go test -v -coverprofile /tmp/$(SERVICE).cover.out ./...
	go tool cover -html /tmp/$(SERVICE).cover.out -o /tmp/$(SERVICE).cover.html
	echo html file path: /tmp/$(SERVICE).cover.html
