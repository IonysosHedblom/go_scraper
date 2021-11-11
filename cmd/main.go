package main

import (
	"log"

	"github.com/ionysoshedblom/go_scraper/internal/api"
	app "github.com/ionysoshedblom/go_scraper/internal/application"
	"github.com/ionysoshedblom/go_scraper/internal/application/scraper"
	"github.com/ionysoshedblom/go_scraper/internal/db"
	"github.com/ionysoshedblom/go_scraper/internal/repository"
)

func main() {
	dsn := db.CreateDbDSN()
	db, err := db.NewDB("postgres", dsn)

	repo := repository.NewRepository(db)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	scraper := scraper.New()
	application := app.NewApplication(scraper, repo)

	httpServer := api.NewApi(application)
	httpServer.Run()
}
