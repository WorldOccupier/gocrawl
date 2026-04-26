package crawl

import (
	"time"

	"com.gocrawl/queue"
)

func StartCrawl() error {
	queue.InitUrls()
	for {
		links, err := queue.GetUrlsToCrawl()
		if err != nil {
			return err
		}

		for _, url := range links {
			pageContent, _ := GetPageContent(url, true)
			pageLinks := GetPageLinks(pageContent)
			queue.AppendUrlsToCrawl(pageLinks)
		}

		time.Sleep(time.Second * 2)
	}
}
