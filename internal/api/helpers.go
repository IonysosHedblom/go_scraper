package api

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func (s *api) CallSource(url string) (*html.Node, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, res.Status)
	}

	doc, err := html.Parse(res.Body)

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return doc, nil
}

func buildUrlWithIngredientsQuery(ingredients []string) string {
	var queryString string = "https://www.ica.se/Templates/ajaxresponse.aspx?ajaxFunction=RecipeListMdsa&num=20&sortbymetadata=Relevance&id=12&mdsarowentityid=ca2947b2-0c0b-4936-b300-a42700eb2734"

	for _, ingredient := range ingredients {
		queryString += fmt.Sprintf("&filter=Ingrediens%%3A%v", strings.Title(ingredient))
	}

	return queryString
}

func buildQueryUrl(query string) string {
	query = strings.ReplaceAll(query, " ", "+")
	url := fmt.Sprintf("https://www.ica.se/Templates/ajaxresponse.aspx?ajaxFunction=RecipeListMdsa&mdsarowentityid=&num=16&query=%s&sortbymetadata=Relevance&id=12", query)
	return url
}

func buildRecipePageUrl(title string, id string) string {
	title = strings.ReplaceAll(title, " ", "-")
	title = strings.ReplaceAll(title, "ä", "a")
	title = strings.ReplaceAll(title, "ö", "o")
	title = strings.ReplaceAll(title, "å", "a")

	url := fmt.Sprintf("https://www.ica.se/recept/%s-%s", title, id)
	fmt.Println(url)
	return url
}
