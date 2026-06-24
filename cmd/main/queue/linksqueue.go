package queue

import (
	"context"
	"math"

	"com.gocrawl/common"
	"com.gocrawl/logger"
	"golang.org/x/time/rate"
)

var (
	ctx               = context.Background()
	initUrlsFile      = "initurls.txt"
	crawlableQueueKey = "queue"
	blmPopDirection   = "LEFT"
	redisClient       = GetRedisClient()
)

func GetUrlsToCrawl(urlsCount rate.Limit) ([]string, error) {
	queueLength, _ := redisClient.LLen(ctx, crawlableQueueKey).Result()
	if queueLength == 0 {
		return []string{}, nil
	}
	_, results, err := redisClient.BLMPop(ctx, 0, blmPopDirection, int64(math.Min(float64(urlsCount), float64(queueLength))), crawlableQueueKey).Result()
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return results, nil
}

func AppendUrlsToCrawl(urls []string) {
	for _, url := range urls {
		redisClient.RPush(ctx, crawlableQueueKey, url)
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
