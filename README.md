# gocrawl

A concurrent web crawler built in Go. Crawls web pages starting from seed URLs, respects robots.txt, extracts links, and stores page content in PostgreSQL — all coordinated via Redis.

## Architecture

```
initurls.txt --> Redis Queue (list) --> doCrawl() goroutines (up to 50)
                    ^                       |
                    |                  +----+-------+
                    |                  |            |
                    |            robots.txt?   Fetch page
                    |                  |            |
                    |                  |     Extract links
                    |                  |     (resolve + fragment strip)
                    |                  |            |
                    |                  |     Save to PostgreSQL
                    |                  |            |
                    +------------------+     Mark visited in Redis (TTL: 7d)
```

## Key details

- Concurrent crawling with rate limiting and robots.txt compliance.
- URL fragments (`#section`) are stripped during link resolution to avoid redundant crawls of the same page.
- Redis is used for both the crawl frontier (queue) and visited URL deduplication.
- PostgreSQL stores crawled page content with upsert-based deduplication.

## Prerequisites

- [Docker](https://docs.docker.com/engine/install/) (for the containerized setup)
- [Go 1.26+](https://go.dev/dl/) (for local builds only)

## Quick start (Docker)

```bash
# Clone and enter the repo
git clone <repo-url> && cd gocrawl

# Edit seed URLs (optional)
# vi cmd/main/initurls.txt

# Start everything (Redis + Postgres + app)
make up

# Follow logs
make logs

# Stop
make down
```

The crawler will start processing URLs and storing results in PostgreSQL automatically.

## Build & run locally

```bash
# Ensure Redis and Postgres are running (Docker)
make redis postgres

# Build the binary
make localbuild
# or: go build -C cmd/main -o ../../gocrawl .

# Run
./gocrawl
```

## Configuration

All configuration is via environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `REDIS_HOST` | `localhost` | Redis server host |
| `REDIS_PORT` | `6379` | Redis server port |
| `DATABASE_URL` | `postgres://user:pass@localhost:5432/mydb` | PostgreSQL connection string |

## Database schema

```sql
CREATE TABLE t_web_page_details (
    url        VARCHAR(2048) PRIMARY KEY,
    crawled_at TIMESTAMP NOT NULL,
    content    TEXT NOT NULL,
    processed  BOOLEAN NOT NULL DEFAULT FALSE
);
```

## Cleanup

```bash
# Stop and remove containers + volumes (wipes all data)
make clean
# or: docker compose down -v

# Remove local binary
rm -f gocrawl
```

## Makefile reference

| Target | Description |
|--------|-------------|
| `up` | Start all services in detached mode |
| `down` | Stop all services |
| `rebuild` | Full rebuild: wipe volumes, rebuild images, start, tail logs |
| `logs` | Tail container logs |
| `clean` | Stop and remove containers + volumes |
| `redis` | Start only Redis |
| `postgres` | Start only Postgres |
| `localbuild` | Build the Go binary locally (output: `./gocrawl`) |
