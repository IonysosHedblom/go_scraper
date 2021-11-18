package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func (s *api) RecipeDetailsRouter(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		s.GetRecipeDetails(w, req)
	default:
		http.Error(w, "No support for this method", http.StatusMethodNotAllowed)
		return
	}
}

func (s *api) TrimHtmlAndCallSrc(url string) (*html.Node, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, res.Status)
	}

	defer res.Body.Close()

	src, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("reading %s: %s", url, res.Status)
	}

	strToReplace := `<li style="display:;"><a href="/handla/" class="navigation__item">Handla online</a> <!----></li>`

	cleanHTML := strings.ReplaceAll(string(src), strToReplace, "")
	cleanerHTML := strings.Split(cleanHTML, `<div class="comment-section__wrapper extra-padding"><div class="comment-section__write-comment"><div aria-label="Skriv din kommentar" class="input-textarea input-textarea--56 input-textarea--simple">`)
	cleanestHTML := cleanerHTML[0] + "</body>" + "\n" + "</html>"
	doc, err := html.Parse(strings.NewReader(cleanestHTML))

	if err != nil {
		return nil, err
	}

	return doc, nil
}

func (s *api) GetRecipeDetails(w http.ResponseWriter, req *http.Request) {
	queries := req.URL.Query()

	if len(queries) > 2 {
		http.Error(w, "bad query param", http.StatusBadRequest)
		return
	}

	recipeTitle := queries["title"][0]
	recipeId := queries["id"][0]

	url := buildRecipePageUrl(recipeTitle, recipeId)
	document, err := s.TrimHtmlAndCallSrc(url)

	if err != nil {
		http.Error(w, "error getting the external source html", http.StatusInternalServerError)
		return
	}

	recipeDetails := s.app.CallRecipeDetailsScraping(document)
	trimmedRecipeDetails := strings.ReplaceAll(recipeDetails, `\`, "")
	fmt.Println(trimmedRecipeDetails)
	// recipeIdAsInt64, err := shared.ConvertStringToInt64(recipeId)

	if err != nil {
		http.Error(w, "cant convert string to int64", http.StatusInternalServerError)
		return
	}

	// err = s.handlers.RecipeHandler.UpdateIngredientsAndChecklist(recipeDetails.Ingredients, recipeDetails.Checklist, *recipeIdAsInt64)

	if err != nil {
		http.Error(w, "db error updating ingredients and checklist", http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, "error marshaling json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(trimmedRecipeDetails))

}
