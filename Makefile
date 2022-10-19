docker_local:
	@echo Starting local docker compose
	docker-compose -f docker-compose-local.yaml up -d --build --remove-orphans

docker_local_down:
	@echo Stoping local docker compose
	docker-compose -f docker-compose-local.yaml down --remove-orphans

docker:
	@echo Starting docker compose
	docker-compose -f docker-compose.yaml up --build --remove-orphans

docker_down:
	@echo Stoping docker compose
	docker-compose -f docker-compose.yaml down --remove-orphans

# helm repo add bitnami https://charts.bitnami.com/bitnami
#=====================================================
# kuberneties
k8s_install:
	helm install kafka bitnami/kafka
	helm install microservice deploy/microservice/

k8s_update:
	helm upgrade microservice deploy/microservice/

k8s_uninstall:
	helm uninstall microservice 
	helm uninstall kafka

# postgresql://postgres:postgres@192.168.39.213:30001/microservice?sslmode=disable
# createdb --username=postgres --owner=postgres microservice_finance
# createdb --username=postgres --owner=postgres microservice_message
# migrate -path internal/adapters/migration -database "postgresql://postgres:postgres@192.168.39.213:30001/microservice_finance?sslmode=disable" -verbose up
# migrate -path internal/adapters/migration -database "postgresql://postgres:postgres@192.168.39.213:30001/microservice?sslmode=disable" -verbose up
# migrate -path internal/adapters/migration -database "postgresql://postgres:postgres@192.168.39.213:30001/microservice_message?sslmode=disable" -verbose up
