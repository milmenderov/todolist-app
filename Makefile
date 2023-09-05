.PHONY: dc run test lint

dc:
	docker-compose -f ./.docker/docker-compose.yml up -d


