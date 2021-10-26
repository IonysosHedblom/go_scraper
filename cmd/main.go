package main

import (
	"github.com/ionysoshedblom/go_scraper/internal/application/api"
	"github.com/ionysoshedblom/go_scraper/internal/application/scraper"
)

var BaseUrl string = "https://www.ica.se/Templates/ajaxresponse.aspx?ajaxFunction=RecipeListMdsa&mdsarowentityid=&num=16&query=pasta&sortbymetadata=Relevance&id=12&_hour=7"

func main() {
	scraper := scraper.New()
	
	applicationAPI := api.NewApplication(scraper)
}

