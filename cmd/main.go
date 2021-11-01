package main

import (
	"github.com/ionysoshedblom/go_scraper/internal/api"
	app "github.com/ionysoshedblom/go_scraper/internal/application"
	"github.com/ionysoshedblom/go_scraper/internal/application/scraper"
)

func main() {
	scraper := scraper.New()
	application := app.NewApplication(scraper)

	httpServer := api.NewApi(application)
	httpServer.Run()
}
