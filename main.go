package main

import (
	"flag"
	"fmt"
	"net/http"
	"search_scraper/csv"
	"search_scraper/scraper"
	"search_scraper/web"
)

func main() {
	csvPath := flag.String("csv_path", "./test.csv", "Path to the CSV file (exported google sheet)")
	csvCollumn := flag.Int("csv_column", 2, "Column index where urls are stored")
	searchQuery := flag.String("search_query", "", "Search query")
	searchLinks := flag.Int("search_links", 100, "Search links count")
	searchLinksPerPage := flag.Int("slpp", 10, "Search links per page")
	serverPort := flag.Int("server_port", 8080, "Server port")
	flag.Parse()

	d := csv.GetAllDoaminNamesFromCsv(*csvPath, *csvCollumn)
	fmt.Println(len(d))
	fu := scraper.ScrapeSearch(*searchQuery, *searchLinks, *searchLinksPerPage, d)

	fn := web.GenerateIndexFile(fu)

	for i, url := range fu {
		fmt.Printf("%d:\t%s\n", i, csv.GetDomainFromUrl(url))
	}

	fmt.Printf("Generated index file: %s\nStarting server...\nhttp://localhost:%d\n", fn, *serverPort)

	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(fmt.Sprintf(":%d", *serverPort), nil)
}
