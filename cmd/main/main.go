package main

import (
	"com.gocrawl/crawl"
	"com.gocrawl/queue"
)

func main() {
	queue.EnsureCrawlableUrlsArePresentOnStartup()
	crawl.StartCrawl()
}
