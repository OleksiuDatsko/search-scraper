package utils

import (
	"fmt"
	"log"
	"net/http"
	"search_scraper/src/types"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
)

var ErrBotDetected = fmt.Errorf("bot detected")

func Srcape(gu string) (types.ScrapedPage, error) {
	var scrapedLinks []types.ScrapedLink
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.google.com%s", gu), nil)
	if err != nil {
		log.Printf("Error: %s \n", err)
		return types.ScrapedPage{}, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error: %s \n", err)
		return types.ScrapedPage{}, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("Error: %s \n", err)
		return types.ScrapedPage{}, err

	}

	if len(doc.Find("form#captcha-form").Nodes) > 0 {
		return types.ScrapedPage{}, ErrBotDetected

	}

	doc.Find("div.g").Each(func(i int, result *goquery.Selection) {
		link, _ := result.Find("a").First().Attr("href")

		scrapedLinks = append(scrapedLinks, types.ScrapedLink{
			ID:      uuid.New().String(),
			Title:   result.Find("h3").First().Text(),
			Link:    link,
			Domain:  GetDomainFromURL(link),
			Snipped: result.Find(".VwiC3b").First().Text(),
		})
	})

	pn, _ := doc.Find("a#pnnext").Attr("href")
	pp, _ := doc.Find("a#pnprev").Attr("href")

	return types.ScrapedPage{
		ScrapedLink:  scrapedLinks,
		PrevPageLink: pp,
		NextPageLink: pn,
	}, nil

}
