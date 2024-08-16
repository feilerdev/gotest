.PHONY: up
up:
	docker-compose up --force-recreate --remove-orphans --no-deps --build