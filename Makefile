.PHONY: build run seed-all

build:
	@docker-compose build

run:
	@docker-compose up

down:
	@docker-compose down
