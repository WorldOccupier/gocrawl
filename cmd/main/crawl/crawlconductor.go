package crawl

import (
	"context"
	"time"

	"com.gocrawl/contenthandler"
	"com.gocrawl/logger"
	"com.gocrawl/queue"
	"github.com/redis/go-redis/v9"

	"golang.org/x/time/rate"
)

var (
	crawledValue = "1"
	crawledTTL = time.Hour*24*7
	rateLimit rate.Limit = 50
	burstLimit = 5
	concurrentGoRoutines = 100
)

func doCrawl(url string, redisClient *redis.Client, webPageContentHandler *contenthandler.WebPageContentHandler, limiter *rate.Limiter) error {
	err := limiter.Wait(context.Background())
	if err != nil {
		return err
	}
	cmd := redisClient.Keys(ctx, url)
	result, err := cmd.Result()
	if err != nil {
		return err
	}

	if len(result) == 0 {
		logger.Log.Info("Crawling: " + url)
		pageContent, err := GetPageContent(url, true)
		if err != nil {
			logger.Log.Error("Error while getting page content for url: " + url, "error", err.Error())
		}
		webPageContentHandler.SaveCrawledContent(url, time.Now(), pageContent)
		pageLinks := GetPageLinks(url, pageContent)
		queue.AppendUrlsToCrawl(pageLinks)
		redisClient.Set(ctx, url, crawledValue, crawledTTL)
	}

	return nil
}

func StartCrawl() error {
	webPageContentHandler := &contenthandler.WebPageContentHandler{}
	redisClient := queue.GetRedisClient()
	limiter := rate.NewLimiter(rateLimit, burstLimit)
	sem := make(chan struct{}, concurrentGoRoutines)

	for {
		links, err := queue.GetUrlsToCrawl(rateLimit)
		if err != nil {
			return err
		}

		for _, url := range links {
			sem <- struct{}{}
			go func(url string) {
				defer func() {
					<-sem
				}()

				doCrawl(url, redisClient, webPageContentHandler, limiter)
			}(url)
		}
	}
}
