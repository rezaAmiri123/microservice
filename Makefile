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
