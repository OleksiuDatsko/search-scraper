package utils

import "regexp"

func GetDomainFromURL(url string) string {
	if url == "" {
		return ""
	}

	patern := `(https?://)?(?P<doamin>([\w\.-]+\.?)+)(/(.+))?`
	re := regexp.MustCompile(patern)
	matches := re.FindStringSubmatch(url)
	return matches[2]
}
