package crawl

import (
	"time"

	"com.gocrawl/contenthandler"
	"com.gocrawl/logger"
	"com.gocrawl/queue"
)

var (
	crawledValue = "1"
	crawledTTL = time.Hour*24*7
)

func StartCrawl() error {
	webPageContentHandler := &contenthandler.WebPageContentHandler{}
	redisClient := queue.GetRedisClient()

	for {
		links, err := queue.GetUrlsToCrawl()
		if err != nil {
			return err
		}

		for _, url := range links {
			cmd := redisClient.Keys(ctx, url)
			result, err := cmd.Result()
			if err != nil {
				return err
			}

			if len(result) == 0 {
				logger.Log.Info("Crawling: " + url)
				pageContent, _ := GetPageContent(url, true)
				webPageContentHandler.SaveCrawledContent(url, time.Now(), pageContent)
				pageLinks := GetPageLinks(url, pageContent)
				queue.AppendUrlsToCrawl(pageLinks)
				redisClient.Set(ctx, url, crawledValue, crawledTTL)
			}
		}

		time.Sleep(time.Millisecond * 200)
	}
}
