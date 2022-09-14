docker_local:
	@echo Starting local docker compose
	docker-compose -f docker-compose-local.yaml up -d --build --remove-orphans

docker_local_down:
	@echo Stoping local docker compose
	docker-compose -f docker-compose-local.yaml down --remove-orphans
