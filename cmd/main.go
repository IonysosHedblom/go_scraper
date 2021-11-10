package main

import (
	"log"

	"github.com/ionysoshedblom/go_scraper/internal/api"
	app "github.com/ionysoshedblom/go_scraper/internal/application"
	"github.com/ionysoshedblom/go_scraper/internal/application/scraper"
	"github.com/ionysoshedblom/go_scraper/internal/repository"
	"github.com/ionysoshedblom/go_scraper/internal/repository/config"
)

func main() {
	dsn := config.CreateDbDSN()
	repo, err := repository.NewRepository("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	defer repo.Close()

	scraper := scraper.New()
	application := app.NewApplication(scraper, repo)

	httpServer := api.NewApi(application)
	httpServer.Run()
}
