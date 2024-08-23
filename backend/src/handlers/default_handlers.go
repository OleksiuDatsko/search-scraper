package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"search_scraper/src/storage"
	"search_scraper/src/types"
	"strconv"
)

func GetLinkslist(st *storage.Storage, listType string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		wl, err := st.GetList(listType)
		if err != nil {
			log.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json_wl, err := json.Marshal(wl)
		if err != nil {
			log.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(json_wl)
	}
}

func PostLinkToLinkslist(st *storage.Storage, listType string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var link types.Link
		json.NewDecoder(r.Body).Decode(&link)
		err := st.AddLinkToList(listType, link)
		if err != nil {
			log.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func GetLinkslistLink(st *storage.Storage, listType string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		link, err := st.GetLinkFromList(listType, id)
		if err != nil {

			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			log.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json_link, err := json.Marshal(link)
		if err != nil {
			log.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(json_link)
	}
}

func DeleteLinkslistLink(st *storage.Storage, listType string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = st.DeleteLinkFromList(listType, id)
		if err != nil {
			log.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func PutLinkslistLink(st *storage.Storage, listType string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var link types.Link
		json.NewDecoder(r.Body).Decode(&link)
		link.ID = id
		err = st.UpdateLinkInList(listType, link)
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			log.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func OptionsHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
	}
}
