package main

import (
	"log"

	"github.com/ionysoshedblom/go_scraper/internal/api"
	"github.com/ionysoshedblom/go_scraper/internal/application"
	"github.com/ionysoshedblom/go_scraper/internal/application/scraper"
	"github.com/ionysoshedblom/go_scraper/internal/db"
	"github.com/ionysoshedblom/go_scraper/internal/repository"
)

func main() {
	dsn := db.CreateDbDSN()
	db, err := db.NewDB("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repo := repository.NewRepository(db)
	scraper := scraper.New()
	app := application.NewApplication(scraper, repo)

	httpServer := api.NewApi(app)
	httpServer.Run()
}
