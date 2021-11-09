package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

func (s *api) ScraperRouter(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		s.GetByQuery(w, req)
	case "POST":
		s.PostWithIngredients(w, req)
	default:
		return
	}
}

func (s *api) GetByQuery(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return
	}

	queries := req.URL.Query()
	if len(queries) > 1 {
		http.Error(w, "too many queries", http.StatusBadRequest)
		return
	}

	q := queries["query"]
	if len(q) > 1 {
		http.Error(w, "too many queries", http.StatusBadRequest)
		return
	}

	url := buildQueryUrl(q[0])

	document, err := s.CallSource(url)

	if err != nil {
		http.Error(w, "bad payload", http.StatusBadRequest)
		return
	}

	response := s.app.CallRecipeResultScraping(document)

	j, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func (s *api) PostWithIngredients(w http.ResponseWriter, req *http.Request) {
	var i entity.Ingredients
	if req.Method != "POST" {
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer req.Body.Close()

	if req.Body == nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	json.Unmarshal(body, &i)
	url := buildUrlWithIngredientsQuery(i.Ingredients)
	document, err := s.CallSource(url)

	if err != nil {
		http.Error(w, "bad payload", http.StatusBadRequest)
	}

	response := s.app.CallRecipeResultScraping(document)
	j, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

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
	var queryString string = "https://www.ica.se/Templates/ajaxresponse.aspx?ajaxFunction=RecipeListMdsa&num=20&sortbymetadata=Relevance&id=12&_hour=11&mdsarowentityid=ca2947b2-0c0b-4936-b300-a42700eb2734"

	for _, ingredient := range ingredients {
		queryString += fmt.Sprintf("&filter=Ingrediens%%3A%v", strings.Title(ingredient))
	}
	fmt.Println(queryString)
	return queryString
}

func buildQueryUrl(query string) string {
	url := fmt.Sprintf("https://www.ica.se/Templates/ajaxresponse.aspx?ajaxFunction=RecipeListMdsa&mdsarowentityid=&num=16&query=%s&sortbymetadata=Relevance&id=12", query)
	return url
}
