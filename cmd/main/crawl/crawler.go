package crawl

import (
	"errors"
	"io"
	"net/http"

	"com.gocrawl/logger"
)

func GetPageContent(url string, checkCrawlability bool) (string, error) {
	if checkCrawlability {
		canCrawl, err := CanCrawl(url)
		if err != nil || canCrawl == false {
			return "", errors.New("Cannot crawl url: " + url)
		}
	}

	response, err := http.Get(url)
	if err != nil {
		logger.Log.Error("Error retrieving content from url: ", "error", err)
		return "", errors.New("Error retrieving content from url")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Log.Error("Error reading response body: ", "error", err.Error())
		return "", errors.New("Error reading response body")
	}

	return string(body), nil
}
