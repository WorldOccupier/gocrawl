package queue

import "github.com/redis/go-redis/v9"

var (
	redisHost = "localhost"
	redisPort = "6379"
)

func GetRedisClient() *redis.Client {
	redisOptions := &redis.Options{
		Addr: redisHost + ":" + redisPort,
	}

	return redis.NewClient(redisOptions)
}
