package main

import (
	"fmt"
	"search_scraper/csv"
	"search_scraper/scraper"

)

func main() {
	domains := csv.GetAllDoaminNamesFromCsv("./test.csv", 2)
	fmt.Println(len(domains))
	scraper.ScrapeSearch("ремонт+під+ключ+львів", 30, domains)
}
