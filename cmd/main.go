package main

import (
	"github.com/ionysoshedblom/go_scraper/internal/application/api"
	"github.com/ionysoshedblom/go_scraper/internal/application/scraper"
	server "github.com/ionysoshedblom/go_scraper/internal/framework/http"
)

func main() {
	scraper := scraper.New()
	applicationAPI := api.NewApplication(scraper)
	
	httpServer := server.NewServer(applicationAPI)
	httpServer.Run()
}
