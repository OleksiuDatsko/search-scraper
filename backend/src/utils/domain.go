package utils

import (
	"regexp"
	"search_scraper/src/types"
)

func GetDomainFromURL(url string) string {
	if url == "" {
		return ""
	}

	patern := `(https?://)?(?P<doamin>([\w\.-]+\.?)+)(/(.+))?`
	re := regexp.MustCompile(patern)
	matches := re.FindStringSubmatch(url)
	return matches[2]
}

func GetLinkFromCVSRow(r []string) types.Link {
	return types.Link{
		Domain:     GetDomainFromURL(r[2]),
		Url:        r[2],
		FilterType: "domain",
	}
}
