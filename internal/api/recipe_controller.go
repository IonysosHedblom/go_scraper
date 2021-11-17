package api

import (
	"encoding/json"
	"net/http"

	"github.com/ionysoshedblom/go_scraper/internal/shared"
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

func (s *api) GetRecipeDetails(w http.ResponseWriter, req *http.Request) {
	queries := req.URL.Query()

	if len(queries) > 2 {
		http.Error(w, "bad query param", http.StatusBadRequest)
		return
	}

	recipeTitle := queries["title"][0]
	recipeId := queries["id"][0]

	url := buildRecipePageUrl(recipeTitle, recipeId)
	document, err := s.CallSource(url)

	if err != nil {
		http.Error(w, "error getting the external source html", http.StatusInternalServerError)
		return
	}

	recipeDetails := s.app.CallRecipeDetailsScraping(document)

	recipeAsInt64, err := shared.ConvertStringToInt64(recipeId)

	if err != nil {
		http.Error(w, "cant convert string to int64", http.StatusInternalServerError)
		return
	}

	err = s.handlers.RecipeHandler.UpdateIngredientsAndChecklist(recipeDetails.Ingredients, recipeDetails.Checklist, *recipeAsInt64)

	if err != nil {
		http.Error(w, "db error updating ingredients and checklist", http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(recipeTitle)

	if err != nil {
		http.Error(w, "error marshaling json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}
