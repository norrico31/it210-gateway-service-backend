build:
	@docker-compose build

run-dev:
	@go run ./cmd

run-docker-up:
	@docker-compose up -d --build

down:
	@docker-compose down -v
