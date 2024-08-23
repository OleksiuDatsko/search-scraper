package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"search_scraper/src/storage"
	"search_scraper/src/utils"

	"strconv"
	"strings"
)

func Search(st *storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		d := r.URL.Query().Get("d")
		generatePage := r.URL.Query().Get("gp")

		if q == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if d == "" {
			d = "0"
		}
		if d == "full" {
			d = "40"
		}

		di, err := strconv.Atoi(d)
		if err != nil {
			log.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		q = strings.ReplaceAll(q, " ", "+")

		res, err := st.FilteredScraping(q, di)
		if err != nil {
			if err == utils.ErrBotDetected {
				w.WriteHeader(http.StatusIMUsed)
				return
			} else {
				log.Printf("Error: %s \n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		if generatePage == "1" {
			w.Header().Add("Content-Type", "text/html")
			w.Write([]byte(GeneratePage(res)))
			return
		}
		json_res, err := json.Marshal(res)
		if err != nil {
			log.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(json_res)
	}
}

func GeneratePage(sr storage.ScrapedResult) string {
	var links string
	for _, link := range sr.ScrapedLinks {
		links += fmt.Sprintf("<li><a href=\"%s\" target=\"_blank\">%s</a></li>\n", link.Link, link.Domain)
	}

	return fmt.Sprintf("<html><body><ol>\n%s</ol></body></html>", links)
}
