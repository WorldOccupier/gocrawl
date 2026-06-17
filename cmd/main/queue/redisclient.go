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
	count, err := redisClient.LLen(ctx, queueKey).Result()
	if err != nil || count == 0 {
		InitUrls()
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
