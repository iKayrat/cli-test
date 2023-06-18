.phony: run up down

run:
	go run cli.go $(arg1) $(arg2)

up:
	docker compose up
down:
	docker compose down

