.PHONY: up down rebuild logs clean

up:
	docker compose up -d

down:
	docker compose down

rebuild:
	docker compose down -v && docker compose up -d --build && docker compose logs -f

logs:
	docker compose logs -f

clean:
	docker compose down -v
