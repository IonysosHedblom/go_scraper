package main

import (
	"github.com/ionysoshedblom/go_scraper/handler"
	"github.com/ionysoshedblom/go_scraper/service"
)

var BaseUrl string = "https://www.ica.se/Templates/ajaxresponse.aspx?ajaxFunction=RecipeListMdsa&mdsarowentityid=&num=16&query=pasta&sortbymetadata=Relevance&id=12&_hour=7"

func main() {
	scraperService := service.NewScraperService()

	handler.NewMicroService(scraperService)
}

