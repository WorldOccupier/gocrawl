package contenthandler

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

var (
	ctx = context.Background()
	defaultDatabaseUrl = "postgres://user:pass@localhost:5432/mydb"
)

type WebPageContentHandler struct {
	connection *pgx.Conn
}

func getDBUrl() string {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = defaultDatabaseUrl
	}

	return databaseUrl
}

func (webPageContentHandler *WebPageContentHandler) GetConnection() *pgx.Conn {
	if webPageContentHandler.connection == nil {
		dbUrl := getDBUrl()
		pgConnection, err := pgx.Connect(ctx, dbUrl)
		if err != nil {
			panic(err)
		}

		webPageContentHandler.connection = pgConnection
	}

	return webPageContentHandler.connection
}

func (webPageContentHandler *WebPageContentHandler) SaveCrawledContent(url string, crawledDateTime time.Time, content string) {
	connection := webPageContentHandler.GetConnection()
	_, err := connection.Exec(ctx, "INSERT INTO t_web_page_details VALUES ($1, $2, $3)", url, crawledDateTime, content)
	if err != nil {
		panic(err)
	}
}
