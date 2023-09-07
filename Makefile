.PHONY: re down run

re:
	docker-compose -f ./.docker/docker-compose.yml up --build -d
run:
	docker-compose -f ./.docker/docker-compose.yml up -d
down:
	docker-compose -f ./.docker/docker-compose.yml down -v

logs:
	docker-compose -f ./.docker/docker-compose.yml logs -f
