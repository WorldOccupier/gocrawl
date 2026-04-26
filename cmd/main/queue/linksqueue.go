package queue

import (
	"context"

	"com.gocrawl/common"
	"com.gocrawl/logger"
)

var (
	ctx          = context.Background()
	initUrlsFile = "initurls.txt"
	queueKey     = "queue"
	redisClient  = GetRedisClient()
)

func GetUrlsToCrawl() ([]string, error) {
	result, err := redisClient.BLPop(ctx, 0, queueKey).Result()
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	urls := make([]string, 0)
	urls = append(urls, result[1])

	return urls, nil
}

func AppendUrlsToCrawl(urls []string) {
	for _, url := range urls {
		redisClient.RPush(ctx, queueKey, url)
	}
}

func InitUrls() {
	urls, err := common.GetFileLines(initUrlsFile)

	if err != nil {
		logger.Log.Error("Error retrieveing links for crawling: ", "error", err.Error())
		return
	}

	AppendUrlsToCrawl(urls)
}
