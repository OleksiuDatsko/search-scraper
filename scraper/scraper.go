package scraper

import (
	"fmt"
	"log"
	"net/http"
	"search_scraper/csv"
	"slices"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// GetData retrieves search result links from Google based on the search query and page number.
//
// Parameters:
// - q: the search query.
// - p: the page number to start the search.
// - fd: a slice of strings containing domain names to exclude.
// - u: a pointer to a slice of strings to store the result links.
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
		domain := csv.GetDomainFromUrl(link)
		if domain == "www.instagram.com" || domain == "www.facebook.com" || !slices.Contains(fd, domain) {
			*u = append(*u, link)
		}
	})
}

// ScrapeSearch performs a search on a website and retrieves URLs based on the given query, page count, and search interval.
//
// Parameters:
// - q: the search query.
// - p: the total number of pages to scrape.
// - s: the number of links at one page. (default: 10)
// - fd: a slice of strings containing domain names that alredy exist in the docs.
//
// Return type: None.
func ScrapeSearch(q string, p int, s int, fd []string) []string {
	var links []string
	var wg sync.WaitGroup

	for i := 0; i < p; i += s {
		wg.Add(1)
		go func(q string, p int, fd []string, u *[]string) {
			defer wg.Done()
			GetData(q, i, fd, &links)
		}(q, i, fd, &links)
	}
	wg.Wait()

	return links
}
