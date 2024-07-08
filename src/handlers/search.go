package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"search_scraper/src/storage"
	"search_scraper/src/utils"
	"strings"
)

func Search(st *storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.RawQuery

		fmt.Println(q)

		if q == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		q = strings.ReplaceAll(q, " ", "+")

		fmt.Printf("Searching for: %s \n", q)
		res := utils.Scraper(q)
		json_res, err := json.Marshal(res)
		if err != nil {
			fmt.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(json_res)
	}
}
