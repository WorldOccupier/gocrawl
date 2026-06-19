package contenthandler

import (
	"context"
	"os"
	"time"

	"com.gocrawl/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ctx = context.Background()
	defaultDatabaseUrl = "postgres://user:pass@localhost:5432/mydb"
)

type WebPageContentHandler struct {
	pgConnectionPool *pgxpool.Pool
}

func getDBUrl() string {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = defaultDatabaseUrl
	}

	return databaseUrl
}

func (webPageContentHandler *WebPageContentHandler) GetConnection() *pgxpool.Pool {
	if webPageContentHandler.pgConnectionPool == nil {
		dbUrl := getDBUrl()
		pgPool, err := pgxpool.New(ctx, dbUrl)
		if err != nil {
			logger.Log.Error("Could not get connection pool")
		}

		webPageContentHandler.pgConnectionPool = pgPool
	}

	return webPageContentHandler.pgConnectionPool
}

func (webPageContentHandler *WebPageContentHandler) SaveCrawledContent(url string, crawledDateTime time.Time, content string) {
	connection := webPageContentHandler.GetConnection()
	_, err := connection.Exec(ctx, "INSERT INTO t_web_page_details (url, crawled_at, content) VALUES ($1, $2, $3) ON CONFLICT (url) DO UPDATE SET crawled_at = $2, content = $3", url, crawledDateTime, content)
	if err != nil {
		logger.Log.Error("Failed to save crawled content", "url", url, "error", err)
	}
}
