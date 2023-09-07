.PHONY: dc down run

build:
	docker-compose -f ./.docker/docker-compose.yml up --build -d
run:
	docker-compose -f ./.docker/docker-compose.yml up
down:
	docker-compose -f ./.docker/docker-compose.yml down -v