package main

import (
	"github.com/ionysoshedblom/go_scraper/internal/application/api"
	"github.com/ionysoshedblom/go_scraper/internal/application/scraper"
	server "github.com/ionysoshedblom/go_scraper/internal/framework/http"
)

// var baseUrl string = "https://www.ica.se/Templates/ajaxresponse.aspx?ajaxFunction=RecipeListMdsa&mdsarowentityid=&num=16&query=pasta&sortbymetadata=Relevance&id=12&_hour=7"

func main() {
	scraper := scraper.New()
	applicationAPI := api.NewApplication(scraper)
	
	httpServer := server.NewServer(applicationAPI)
	httpServer.Run()
	// src, err := applicationAPI.GetSource(baseUrl)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// htmlSrc, err := applicationAPI.HandleSource(src)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(htmlSrc)
}
