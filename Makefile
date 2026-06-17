.PHONY: up down rebuild logs clean

up:
	docker compose up -d

down:
	docker compose down

rebuild:
	docker compose down -v && docker compose up --build

logs:
	docker compose logs -f

clean:
	docker compose down -v
