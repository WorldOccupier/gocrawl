package main

import (
	"log"

	"com.gocrawl/crawl"
	"com.gocrawl/queue"
)

func main() {
	queue.EnsureCrawlableUrlsArePresentOnStartup()
	err := crawl.StartCrawl()
	if err != nil {
		log.Fatal(err)
	}
}
