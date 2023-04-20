docker_local:
	@echo Starting local docker compose
	docker-compose -f docker-compose-local.yaml up -d --build --remove-orphans

docker_local_down:
	@echo Stoping local docker compose
	docker-compose -f docker-compose-local.yaml down --remove-orphans

docker_mallbots:
	@echo Starting mallbots docker compose
	docker-compose -f docker-compose-mallbots.yaml up --remove-orphans
	# docker compose--profile monolith up

docker_mallbots_down:
	@echo Stoping mallbots docker compose
	docker-compose -f docker-compose-mallbots.yaml down --remove-orphans

docker:
	@echo Starting docker compose
	docker-compose -f docker-compose.yaml up --build --remove-orphans

docker_down:
	@echo Stoping docker compose
	docker-compose -f docker-compose.yaml down --remove-orphans

rebuild-users:
	docker image rm -f mallbots-users
	docker build -t mallbots-users --file docker/Dockerfile.microservices --build-arg=service=users .

build-users:
	docker build -t mallbots-users --file docker/Dockerfile.microservices --build-arg=service=users .

rebuild-baskets:
	docker image rm -f mallbots-baskets
	docker build -t mallbots-baskets --file docker/Dockerfile.microservices --build-arg=service=baskets .
build-baskets:
	docker build -t mallbots-baskets --file docker/Dockerfile.microservices --build-arg=service=baskets .

rebuild-stores:
	docker image rm -f mallbots-stores
	docker build -t mallbots-stores --file docker/Dockerfile.microservices --build-arg=service=stores .
build-stores:
	docker build -t mallbots-stores --file docker/Dockerfile.microservices --build-arg=service=stores .

rebuild-search:
	docker image rm -f mallbots-search
	docker build -t mallbots-search --file docker/Dockerfile.microservices --build-arg=service=search .
build-search:
	docker build -t mallbots-search --file docker/Dockerfile.microservices --build-arg=service=search .

rebuild-payments:
	docker image rm -f mallbots-payments
	docker build -t mallbots-payments --file docker/Dockerfile.microservices --build-arg=service=payments .
build-payments:
	docker build -t mallbots-payments --file docker/Dockerfile.microservices --build-arg=service=payments .

rebuild-ordering:
	docker image rm -f mallbots-ordering
	docker build -t mallbots-ordering --file docker/Dockerfile.microservices --build-arg=service=ordering .
build-ordering:
	docker build -t mallbots-ordering --file docker/Dockerfile.microservices --build-arg=service=ordering .

rebuild-depot:
	docker image rm -f mallbots-depot
	docker build -t mallbots-depot --file docker/Dockerfile.microservices --build-arg=service=depot .
build-depot:
	docker build -t mallbots-depot --file docker/Dockerfile.microservices --build-arg=service=depot .

rebuild-cosec:
	docker image rm -f mallbots-cosec
	docker build -t mallbots-cosec --file docker/Dockerfile.microservices --build-arg=service=cosec .
build-cosec:
	docker build -t mallbots-cosec --file docker/Dockerfile.microservices --build-arg=service=cosec .

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
#
#
# //go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination interface_mock.go -self_package github.com/uber/cadence/client/frontend
#
