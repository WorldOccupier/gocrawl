package crawl

import (
	"errors"
	"io"
	"net/http"

	"com.gocrawl/logger"
)

var (
	get = "GET"
	userAgent = "user-agent"
	gogetbot = "gogetbot"
)

func makeRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	request, err := http.NewRequest(get, url, nil)
	if err != nil {
		logger.Log.Error("Error while creating request for url: " + url + err.Error())
		return nil, err
	}
	request.Header.Set(userAgent, gogetbot)

	return client.Do(request)
}

func GetPageContent(url string, checkCrawlability bool) (string, error) {
	if checkCrawlability {
		canCrawl, err := CanCrawl(url)
		if err != nil || !canCrawl {
			return "", errors.New("Cannot crawl url: " + url)
		}
	}

	response, err := makeRequest(url)
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
