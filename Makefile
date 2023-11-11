.PHONY: re down run clean db logs

re:
	docker-compose -f ./.docker/docker-compose.yml up --build -d
run:
	docker-compose -f ./.docker/docker-compose.yml up -d
down:
	docker-compose -f ./.docker/docker-compose.yml down -v
db:
	docker-compose -f ./.docker/docker-compose.yml up -d db

clean:
	docker-compose -f ./.docker/docker-compose.yml down -v && sudo rm -rf ./.docker/.database

logs:
	docker-compose -f ./.docker/docker-compose.yml logs -f
