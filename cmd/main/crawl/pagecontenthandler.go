package crawl

import "regexp"

func GetPageLinks(pageContent string) []string {
	re := regexp.MustCompile(`href=["'](https://.*?)["']`)
	matches := re.FindAllStringSubmatch(pageContent, -1)
	links := make([]string, 0)

	for _, m := range matches {
		links = append(links, m[1])
	}

	return links
}
