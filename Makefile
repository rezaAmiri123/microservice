REPO_NAME=reza879
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
	docker-compose -f docker-compose.yaml up --build --remove-orphans --no-cache mallbots-baskets

docker_down:
	@echo Stoping docker compose
	docker-compose -f docker-compose.yaml down --remove-orphans

build-services: build-users build-baskets build-notifications build-stores build-search build-payments build-ordering build-depot build-cosec  

rebuild-users:
	docker image rm -f mallbots-users $(REPO_NAME)/users:latest
	docker build -t mallbots-users -t $(REPO_NAME)/users:latest --file docker/Dockerfile.microservices --build-arg=service=users .
build-users:
	docker build -t mallbots-users -t $(REPO_NAME)/users:latest --file docker/Dockerfile.microservices --build-arg=service=users .
minikube-load-users:
	minikube image load $(REPO_NAME)/users:latest

rebuild-baskets:
	docker image rm -f mallbots-baskets $(REPO_NAME)/baskets:latest
	docker build -t mallbots-baskets -t $(REPO_NAME)/baskets:latest --file docker/Dockerfile.microservices --build-arg=service=baskets .
build-baskets:
	docker build -t mallbots-baskets -t $(REPO_NAME)/baskets:latest --file docker/Dockerfile.microservices --build-arg=service=baskets .
minikube-load-baskets:
	minikube image load $(REPO_NAME)/baskets:latest

rebuild-notifications:
	docker image rm -f mallbots-notifications $(REPO_NAME)/notifications:latest
	docker build -t mallbots-notifications -t $(REPO_NAME)/notifications:latest --file docker/Dockerfile.microservices --build-arg=service=notifications .
build-notifications:
	docker build -t mallbots-notifications -t $(REPO_NAME)/notifications:latest --file docker/Dockerfile.microservices --build-arg=service=notifications .
minikube-load-notifications:
	minikube image load $(REPO_NAME)/notifications:latest

rebuild-stores:
	docker image rm -f mallbots-stores $(REPO_NAME)/stores:latest
	docker build -t mallbots-stores -t $(REPO_NAME)/stores:latest --file docker/Dockerfile.microservices --build-arg=service=stores .
build-stores:
	docker build -t mallbots-stores -t $(REPO_NAME)/stores:latest --file docker/Dockerfile.microservices --build-arg=service=stores .
minikube-load-stores:
	minikube image load $(REPO_NAME)/stores:latest

rebuild-search:
	docker image rm -f mallbots-search $(REPO_NAME)/search:latest
	docker build -t mallbots-search -t $(REPO_NAME)/search:latest --file docker/Dockerfile.microservices --build-arg=service=search .
build-search:
	docker build -t mallbots-search -t $(REPO_NAME)/search:latest --file docker/Dockerfile.microservices --build-arg=service=search .
minikube-load-search:
	minikube image load $(REPO_NAME)/search:latest

rebuild-payments:
	docker image rm -f mallbots-payments $(REPO_NAME)/payments:latest
	docker build -t mallbots-payments -t $(REPO_NAME)/payments:latest --file docker/Dockerfile.microservices --build-arg=service=payments .
build-payments:
	docker build -t mallbots-payments -t $(REPO_NAME)/payments:latest --file docker/Dockerfile.microservices --build-arg=service=payments .
minikube-load-payments:
	minikube image load $(REPO_NAME)/payments:latest

rebuild-ordering:
	docker image rm -f mallbots-ordering $(REPO_NAME)/ordering:latest
	docker build -t mallbots-ordering -t $(REPO_NAME)/ordering:latest --file docker/Dockerfile.microservices --build-arg=service=ordering .
build-ordering:
	docker build -t mallbots-ordering -t $(REPO_NAME)/ordering:latest --file docker/Dockerfile.microservices --build-arg=service=ordering .
minikube-load-ordering:
	minikube image load $(REPO_NAME)/ordering:latest

rebuild-depot:
	docker image rm -f mallbots-depot $(REPO_NAME)/depot:latest
	docker build -t mallbots-depot -t $(REPO_NAME)/depot:latest --file docker/Dockerfile.microservices --build-arg=service=depot .
build-depot:
	docker build -t mallbots-depot -t $(REPO_NAME)/depot:latest --file docker/Dockerfile.microservices --build-arg=service=depot .
minikube-load-depot:
	minikube image load $(REPO_NAME)/depot:latest

rebuild-cosec:
	docker image rm -f mallbots-cosec $(REPO_NAME)/cosec:latest
	docker build -t mallbots-cosec -t $(REPO_NAME)/cosec:latest --file docker/Dockerfile.microservices --build-arg=service=cosec .
build-cosec:
	docker build -t mallbots-cosec -t $(REPO_NAME)/cosec:latest --file docker/Dockerfile.microservices --build-arg=service=cosec .
minikube-load-cosec:
	minikube image load $(REPO_NAME)/cosec:latest

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

helm_install:
	helm install microservice deploy/helm/

helm_update:
	helm upgrade microservice deploy/helm/

helm_update_dependency:
	helm dependency update deploy/helm

helm_uninstall:
	helm uninstall microservice 


minikube_start:
	minikube start --driver=kvm2

minikube_stop:
	minikube stop

minikube_dashboard:
	minikube dashboard

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
