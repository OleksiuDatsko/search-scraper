package handlers

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"search_scraper/src/storage"
	"search_scraper/src/utils"
)

func ImportFindedlist(st *storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving the file", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		defer file.Close()
		reader := csv.NewReader(file)
		err = st.CleanList("findedlist")
		if err != nil {
			fmt.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error reading CSV file", http.StatusInternalServerError)
				return
			}
			
			if record[2] == "" {
				continue
			}
			err = st.AddLinkToList("findedlist", utils.GetLinkFromCVSRow(record))
			if err != nil {
				fmt.Printf("Error: %s \n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		fmt.Fprintf(w, "File imported successfully")
	}
}
