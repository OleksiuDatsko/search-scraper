package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"search_scraper/src/storage"
	"search_scraper/src/types"
	"strconv"
)

func GetWhitelist(st *storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		wl, err := st.GetList("whitelist")
		if err != nil {
			fmt.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json_wl, err := json.Marshal(wl)
		if err != nil {
			fmt.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(json_wl)
	}
}

func PostLinkToWhitelist(st *storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var link types.Link
		json.NewDecoder(r.Body).Decode(&link)
		err := st.AddLinkToList("whitelist", link)
		if err != nil {
			fmt.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func GetWhitelistLink(st *storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		link, err := st.GetLinkFromList("whitelist", id)
		if err != nil {

			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			fmt.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json_link, err := json.Marshal(link)
		if err != nil {
			fmt.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(json_link)
	}
}

func DeleteWhitelistLink(st *storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = st.DeleteLinkFromList("whitelist", id)
		if err != nil {
			fmt.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func PutWhitelistLink(st *storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var link types.Link
		json.NewDecoder(r.Body).Decode(&link)
		link.ID = id
		err = st.UpdateLinkInList("whitelist", link)
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			fmt.Printf("Error: %s \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
