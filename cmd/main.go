package main

import (
	"fmt"

	"github.com/ionysoshedblom/go_scraper/internal/api"
	"github.com/ionysoshedblom/go_scraper/internal/application/scraper"
)

var baseUrl string = "https://www.ica.se/Templates/ajaxresponse.aspx?ajaxFunction=RecipeListMdsa&mdsarowentityid=&num=16&query=pasta&sortbymetadata=Relevance&id=12&_hour=7"

func main() {
	scraper := scraper.New()
	applicationAPI := api.NewApplication(scraper)
	src, err := applicationAPI.GetSource(baseUrl)

	if err != nil {
		fmt.Println(err)
	}

	htmlSrc, err := applicationAPI.HandleSource(src)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(htmlSrc)
}
