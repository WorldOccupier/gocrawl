package queue

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func GetRedisClient() *redis.Client {
	redisHost := getEnv("REDIS_HOST", "localhost")
	redisPort := getEnv("REDIS_PORT", "6379")

	redisOptions := &redis.Options{
		Addr: redisHost + ":" + redisPort,
	}

	return redis.NewClient(redisOptions)
}

func EnsureCrawlableUrlsArePresentOnStartup() {
	links, err := GetUrlsToCrawl()
	if err != nil || len(links) == 0 {
		InitUrls()
	} else if len(links) > 0 {
		AppendUrlsToCrawl(links)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
