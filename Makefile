.PHONY: up down rebuild logs clean localbuild

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

localbuild:
	go build -C cmd/main -o ../../gocrawl .
