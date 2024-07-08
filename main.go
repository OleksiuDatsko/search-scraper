package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"search_scraper/src/handlers"
	"search_scraper/src/middelware"
	"search_scraper/src/storage"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	fmt.Println("Connecting to database...")
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Printf("Error: %s \n", err)
		return
	}
	defer func() {
		db.Close()
		fmt.Println("Closing DB connection")
	}()
	fmt.Println("Connected")
	st, err := storage.Init(db)
	if err != nil {
		fmt.Printf("Error: %s \n", err)
		return
	}

	fmt.Println("Running migrations...")
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
	fmt.Println("Done")

	http.HandleFunc("GET /search", handlers.Search(st))

	http.HandleFunc("GET /whitelist", middelware.LogRequest(handlers.GetLinkslist(st, "whitelist")))
	http.HandleFunc("POST /whitelist", middelware.LogRequest(handlers.PostLinkToLinkslist(st, "whitelist")))
	http.HandleFunc("GET /whitelist/{id}", middelware.LogRequest(handlers.GetLinkslistLink(st, "whitelist")))
	http.HandleFunc("DELETE /whitelist/{id}", middelware.LogRequest(handlers.DeleteLinkslistLink(st, "whitelist")))
	http.HandleFunc("PUT /whitelist/{id}", middelware.LogRequest(handlers.PutLinkslistLink(st, "whitelist")))

	http.HandleFunc("GET /blacklist", middelware.LogRequest(handlers.GetLinkslist(st, "blacklist")))
	http.HandleFunc("POST /blacklist", middelware.LogRequest(handlers.PostLinkToLinkslist(st, "blacklist")))
	http.HandleFunc("GET /blacklist/{id}", middelware.LogRequest(handlers.GetLinkslistLink(st, "blacklist")))
	http.HandleFunc("DELETE /blacklist/{id}", middelware.LogRequest(handlers.DeleteLinkslistLink(st, "blacklist")))
	http.HandleFunc("PUT /blacklist/{id}", middelware.LogRequest(handlers.PutLinkslistLink(st, "blacklist")))

	http.HandleFunc("GET /findedlist", middelware.LogRequest(handlers.GetLinkslist(st, "findedlist")))
	http.HandleFunc("POST /findedlist", middelware.LogRequest(handlers.PostLinkToLinkslist(st, "findedlist")))
	http.HandleFunc("GET /findedlist/{id}", middelware.LogRequest(handlers.GetLinkslistLink(st, "findedlist")))
	http.HandleFunc("PUT /findedlist/{id}", middelware.LogRequest(handlers.PutLinkslistLink(st, "findedlist")))
	http.HandleFunc("DELETE /findedlist/{id}", middelware.LogRequest(handlers.DeleteLinkslistLink(st, "findedlist")))
	http.HandleFunc("POST /findedlist/import", middelware.LogRequest(handlers.ImportFindedlist(st)))

	fmt.Println("Starting server...")
	server := &http.Server{
		Addr: ":8080",
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
