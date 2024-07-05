package web

import (
	"fmt"
	"os"
	"search_scraper/csv"
)

func GenerateIndexFile(l []string) string {
	var links string
	for _, url := range l {
		links += fmt.Sprintf("<li><a href=\"%s\" target=\"_blank\">%s</a></li>\n", url, csv.GetDomainFromUrl(url))
	}

	f, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("<html><body><ol>\n%s</ol></body></html>", links))
	return f.Name()
}
