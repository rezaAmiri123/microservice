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
	# helm install postgres bitnami/postgresql
	helm install microservice deploy/microservice/

k8s_update:
	helm upgrade test-microservice deploy/test-microservice/

k8s_uninstall:
	helm uninstall microservice 
	# helm uninstall postgres 

make_service_user_image:
	cd service_user && docker build -f Dockerfile -t reza879/service_user:$(git rev-parse --short HEAD) . && cd ..
	docker push reza879/service_user:$(git rev-parse --short HEAD)