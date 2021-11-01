package main

import (
	"github.com/ionysoshedblom/go_scraper/internal/api"
	app "github.com/ionysoshedblom/go_scraper/internal/application"
	"github.com/ionysoshedblom/go_scraper/internal/application/scraper"
)

func main() {
	scraper := scraper.New()
	applicationAPI := app.NewApplication(scraper)

	httpServer := api.NewServer(applicationAPI)
	httpServer.Run()
}
