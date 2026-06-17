package crawl

import (
	"net/url"
	"regexp"

	"com.gocrawl/logger"
)

func GetPageLinks(crawledUrl string, pageContent string) []string {
	re := regexp.MustCompile(`href=["'](.*?)["']`)
	matches := re.FindAllStringSubmatch(pageContent, -1)
	links := make([]string, 0)

	parsedUrl, err := url.Parse(crawledUrl)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil
	}

	for _, m := range matches {
		link := m[1]
		if len(link) > 0 && link[0] == '/' {
			link = (&url.URL{Scheme: parsedUrl.Scheme, Host: parsedUrl.Host, Path: link}).String()
		}
		links = append(links, link)
	}

	return links
}
