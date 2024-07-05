package types

type Link struct {
	ID int `json:"id"`
	Url string `json:"url"`
	Domain string `json:"domain"`
	FilterType string `json:"filter_type"`
}