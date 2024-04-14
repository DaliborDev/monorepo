.PHONY: up down clean

up:
	docker-compose up --build --remove-orphans
down:
	docker-compose
clean:
	docker rm -f $(docker ps -aq) && docker rmi $(docker images -q)