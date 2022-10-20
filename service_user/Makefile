DB_NAME=microservice
DB_USER=postgres
DB_PASSWORD=postgres
DB_ADDR=microservice_postgesql
DB_URL=postgresql://$(DB_USER):$(DB_PASSWORD)@$(api_postgesql):5432/$(DB_NAME)?sslmode=disable
MIGRATION_PATH=internal/adapters/migration

.PHONY: gencert
gencert:
	cfssl gencert \
		-initca tls/config/ca-csr.json | cfssljson -bare ca
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=tls/config/ca-config.json \
		-profile=server \
		tls/config/server-csr.json | cfssljson -bare server
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=tls/config/ca-config.json \
		-profile=client \
		tls/config/client-csr.json | cfssljson -bare client
		mv *.pem *.csr tls

####################################
# database
# DB_URL=postgresql://postgres:postgres@192.168.39.213:30001/microservice?sslmode=disable
createdb:
	docker exec -it $(DB_ADDR) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it $(DB_ADDR) dropdb $(DB_NAME)
	
psql:
	psql -h $(DB_ADDR) -U $(DB_USER) -d $(DB_NAME) -W

evans:
	evans --host localhost --port 5080 -r repl
	# evans --host 192.168.39.213 --port 30000 -r repl

# set evanse header
# header Authorization='Bearer xxxxx'

run:
	go run cmd/user/main.go

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


test_coverage:
	go test --cover ./...

user_proto:
	protoc \
		--go_out=proto/grpc --go_opt=paths=source_relative \
		--go-grpc_out=proto/grpc --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=proto/grpc --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=service_user \
		--proto_path=proto/grpc proto/grpc/*.proto

make_image:
	docker build -f Dockerfile -t reza879/service_user:$(git rev-parse --short HEAD) .
	docker push reza879/service_user:$(git rev-parse --short HEAD)