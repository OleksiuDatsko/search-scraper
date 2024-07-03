package scraper

import (
	"fmt"
	"log"
	"net/http"
	"search_scraper/csv"
	"slices"

	"github.com/PuerkitoBio/goquery"
)

type Link struct {
	Href string
	Domain string
}

func GetData(q string, p int, fd []string, u *[]string) {
	url := fmt.Sprintf("https://www.google.com/search?q=%s&start=%d", q, p)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("div.g").Each(func(i int, result *goquery.Selection) {
		link, _ := result.Find("a").First().Attr("href")
		fmt.Printf("%d %s\n", i, link)
		domain := csv.GetDomainFromUrl(link)
		if !slices.Contains(fd, domain) {
			*u = append(*u, link)
		}
	})
}

func ScrapeSearch(q string, p int, fd []string) {
	var urls []string
	for i := 0; i < p; i += 10 {
		GetData(q, i, fd, &urls)
	}
	fmt.Println(urls, len(urls))
}
