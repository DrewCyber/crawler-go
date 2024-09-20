package main

import (
	"database/sql"

	"github.com/DrewCyber/crawler-go/internal/crawler"
	"github.com/DrewCyber/crawler-go/internal/scraper"
	storage "github.com/DrewCyber/crawler-go/internal/sqlite_storage"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	store, err := storage.NewSqliteStore(db)
	if err != nil {
		panic(err)
	}

	scraper := scraper.NewCollector(store)
	crawler := crawler.NewCrawler(scraper, store)

	// open the target URL
	crawler.Start("https://evo-lutio.livejournal.com/1699420.html")
	// collector.Visit("https://evo-lutio.livejournal.com/2024/07/14/")

}
