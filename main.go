package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"search_scraper/src/handlers"
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
	st := storage.Init(db)

	fmt.Println("Running migrations...")
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
	fmt.Println("Done")

	http.HandleFunc("GET /whitelist", handlers.GetWhitelist(st))
	http.HandleFunc("POST /whitelist", handlers.PostLinkToWhitelist(st))
	http.HandleFunc("GET /whitelist/{id}", handlers.GetWhitelistLink(st))
	http.HandleFunc("DELETE /whitelist/{id}", handlers.DeleteWhitelistLink(st))
	http.HandleFunc("PUT /whitelist/{id}", handlers.PutWhitelistLink(st))

	http.HandleFunc("GET /blacklist", handlers.GetBlacklist(st))
	http.HandleFunc("POST /blacklist", handlers.PostLinkToBlacklist(st))
	http.HandleFunc("GET /blacklist/{id}", handlers.GetBlacklistLink(st))
	http.HandleFunc("DELETE /blacklist/{id}", handlers.DeleteBlacklistLink(st))
	http.HandleFunc("PUT /blacklist/{id}", handlers.PutBlacklistLink(st))


	server := &http.Server{
		Addr: ":8080",
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
