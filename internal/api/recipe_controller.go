package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
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
	var recipe entity.Recipe
	queries := req.URL.Query()
	fmt.Println(queries)
	if len(queries) > 2 {
		http.Error(w, "bad query param", http.StatusBadRequest)
		return
	}

	recipeTitle := queries["title"][0]
	recipeId := queries["id"][0]

	recipeIdAsInt64, err := shared.ConvertStringToInt64(recipeId)

	if err != nil {
		http.Error(w, "cant convert string to int64", http.StatusInternalServerError)
		return
	}

	recipeInDb, err := s.handlers.RecipeHandler.GetRecipeById(*recipeIdAsInt64)

	if err != nil {
		http.Error(w, "error fetching recipe from db", http.StatusInternalServerError)
		return
	}

	if recipeInDb.Checklist == nil {
		url := buildRecipePageUrl(recipeTitle, recipeId)
		document, err := s.TrimHtmlAndCallSrc(url)

		if err != nil {
			http.Error(w, "error getting the external source html", http.StatusInternalServerError)
			return
		}
		r := s.app.CallRecipeDetailsScraping(document)
		var response entity.RecipeDetailsResponse

		err = json.Unmarshal([]byte(r), &response)

		if err != nil {
			http.Error(w, "error getting recipedetails in struct", http.StatusInternalServerError)
			return
		}

		recipeDetails := &entity.RecipeDetails{
			Ingredients:  response.RecipeIngredient,
			Instructions: getInstructions(response.RecipeInstructions),
			Rating:       fmt.Sprintf("%f", response.Rating.RatingValue),
		}

		err = s.handlers.RecipeHandler.Update(recipeDetails.Ingredients, recipeDetails.Instructions, recipeDetails.Rating, *recipeIdAsInt64)

		if err != nil {
			http.Error(w, "db error updating ingredients and checklist", http.StatusInternalServerError)
			return
		}

		recipePointer, err := s.handlers.RecipeHandler.GetRecipeById(*recipeIdAsInt64)
		if err != nil {
			http.Error(w, "db error returning recipe", http.StatusInternalServerError)
			return
		}

		recipe = *recipePointer
	} else {
		recipePointer, err := s.handlers.RecipeHandler.GetRecipeById(*recipeIdAsInt64)
		if err != nil {
			http.Error(w, "db error returning recipe", http.StatusInternalServerError)
			return
		}

		recipe = *recipePointer
	}

	j, _ := json.Marshal(recipe)

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
