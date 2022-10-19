DB_NAME=microservice
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
	# evans --host 192.168.39.213 --port 30003 -r repl

# set evanse header
# header Authorization='Bearer xxxxx'

run:
	go run cmd/api/main.go

migratecreate:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq service_user

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

make_image:
	docker build -f Dockerfile -t reza879/service_api:$(git rev-parse --short HEAD) .
	docker push reza879/service_api:$(git rev-parse --short HEAD)