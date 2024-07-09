package types

type Link struct {
	ID         int    `json:"id"`
	Url        string `json:"url"`
	Domain     string `json:"domain"`
	FilterType string `json:"filter_type"`
}

type ScrapedLink struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Domain  string `json:"domain"`
	Snipped string `json:"snipped"`
}

type ScrapedPage struct {
	ScrapedLink  []ScrapedLink `json:"scraped_link"`
	NextPageLink string        `json:"next_page_link"`
	PrevPageLink string        `json:"prev_page_link"`
}
