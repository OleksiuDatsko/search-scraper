package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"search_scraper/src/storage"
	"strconv"
	"strings"
)

func Search(st *storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		d := r.URL.Query().Get("d")

		fmt.Println(q)

		if q == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if d == "" {
			d = "0"
		}

		di, err := strconv.Atoi(d)
		if err != nil {
			fmt.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		q = strings.ReplaceAll(q, " ", "+")

		res := st.FilteredScraping(q, di)
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
