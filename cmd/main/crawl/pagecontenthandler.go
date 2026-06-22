package crawl

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func resolveURL(base *url.URL, href string) *url.URL {
	parsedUrl, err := url.Parse(strings.TrimSpace(href))
	if err != nil {
		return nil
	}
	parsedUrl.Fragment = ""

	if !parsedUrl.IsAbs() {
		parsedUrl = base.ResolveReference(parsedUrl)
	}

	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return nil
	}

	if parsedUrl.String() == base.String() {
		return nil
	}

	return parsedUrl
}

func GetPageLinks(crawledUrl string, pageContent string) []string {
	parsedUrl, _ := url.Parse(crawledUrl)
	var links []string
	reader := strings.NewReader(pageContent)
	tokenizer := html.NewTokenizer(reader)
	for {
		token := tokenizer.Next()
		if token == html.ErrorToken {
			break
		}
		name, hasAttr := tokenizer.TagName()
		if string(name) == "a" && hasAttr {
			for {
				key, val, more := tokenizer.TagAttr()
				if string(key) == "href" {
					link := resolveURL(parsedUrl, string(val))
					if link != nil {
						links = append(links, link.String())
					}
				}

				if !more {
					break
				}
			}
		}
	}

	return links
}
