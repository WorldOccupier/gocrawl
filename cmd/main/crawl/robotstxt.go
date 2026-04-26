package crawl

import (
	"context"
	neturl "net/url"
	"time"

	"com.gocrawl/queue"
	"github.com/jimsmart/grobotstxt"
)

const (
	botUA                    = "gogetbot"
	robotsPath               = "/robots.txt"
	redisPrefix              = "pageCrawlable:"
	crawlabilitySaveDuration = time.Hour * 24 * 7
	cannotCrawlValue         = "0"
	canCrawlValue            = "1"
)

var ctx = context.Background()

func getRobotsUrl(parsedUrl *neturl.URL) string {
	return (&neturl.URL{
		Scheme: parsedUrl.Scheme,
		Host:   parsedUrl.Host,
		Path:   robotsPath}).String()
}

func saveCrawlability(parsedUrl *neturl.URL, canCrawl bool) {
	redisClient := queue.GetRedisClient()
	canCrawlRedisValue := cannotCrawlValue
	if canCrawl {
		canCrawlRedisValue = canCrawlValue
	}
	urlWithoutScheme := parsedUrl.Host + parsedUrl.Path
	redisClient.Set(ctx,
		redisPrefix+urlWithoutScheme,
		canCrawlRedisValue,
		crawlabilitySaveDuration)
}

func CanCrawl(url string) (bool, error) {
	parsedUrl, err := neturl.Parse(url)
	if err != nil {
		return false, err
	}

	robotsUrl := getRobotsUrl(parsedUrl)
	robotsResponse, err := GetPageContent(robotsUrl, false)
	canCrawl := grobotstxt.AgentAllowed(robotsResponse, botUA, url)
	saveCrawlability(parsedUrl, canCrawl)

	return canCrawl, nil
}
