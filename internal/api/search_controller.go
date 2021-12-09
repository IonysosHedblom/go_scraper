package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
)

func (s *api) ScraperRouter(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		s.Get(w, req)
	case "POST":
		s.Post(w, req)
	default:
		http.Error(w, "No support for this method", http.StatusMethodNotAllowed)
		return
	}
}

func (s *api) Get(w http.ResponseWriter, req *http.Request) {
	var recipes []entity.Recipe

	queries := req.URL.Query()
	if len(queries) != 1 {
		http.Error(w, "bad query param", http.StatusBadRequest)
		return
	}

	q := queries["query"]
	if len(q) != 1 {
		http.Error(w, "bad query param", http.StatusBadRequest)
		return
	}

	performedQueryInDb, err := s.handlers.PqHandler.GetPerformedQuery(q[0])

	if err != nil && err.Error() != "sql: no rows in result set" {
		http.Error(w, "error getting performed query from db", http.StatusInternalServerError)
		return
	}

	if performedQueryInDb == nil {
		url := buildQueryUrl(q[0])
		document, err := s.CallSource(url)

		if err != nil {
			http.Error(w, "bad payload", http.StatusBadRequest)
			return
		}

		recipes, err = s.app.CallRecipeResultScraping(document)

		if err != nil {
			http.Error(w, "error scraping recipe ids", http.StatusInternalServerError)
			return
		}

		newQueryId, err := s.handlers.PqHandler.CreateNewPerformedQuery(q[0])
		if err != nil {
			http.Error(w, "error with db conn", http.StatusBadRequest)
			return
		}

		for _, r := range recipes {
			recipeInDb, err := s.handlers.RecipeHandler.GetRecipeById(r.Id)

			if err != nil && err.Error() != "sql: no rows in result set" {
				http.Error(w, "error getting performed query from db", http.StatusInternalServerError)
				return
			}

			if recipeInDb == nil {
				recipe := &entity.Recipe{
					Id:          r.Id,
					Title:       r.Title,
					Description: r.Description,
					ImageUrl:    r.ImageUrl,
					Ingredients: r.Ingredients,
					QueryId:     newQueryId,
				}
				err := s.handlers.RecipeHandler.CreateNewRecipe(recipe)

				if err != nil {
					http.Error(w, "error creating new recipe", http.StatusInternalServerError)
					return
				}
			} else {
				err := s.handlers.RecipeHandler.UpdateRecipeQueryId(newQueryId, r.Id)
				if err != nil {
					http.Error(w, "error updating recipe", http.StatusInternalServerError)
					return
				}
			}
		}
	} else {
		recipes, err = s.handlers.RecipeHandler.GetRecipesByQueryId(int64(performedQueryInDb.Id))
		if err != nil {
			http.Error(w, "error with db conn", http.StatusInternalServerError)
			return
		}
	}

	j, err := json.Marshal(recipes)

	if err != nil {
		http.Error(w, "Error marshaling json response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func (s *api) Post(w http.ResponseWriter, req *http.Request) {
	var recipes []entity.Recipe
	var i entity.ScraperRequest

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer req.Body.Close()

	if req.Body == http.NoBody {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	json.Unmarshal(body, &i)

	ingredientSearchInDb, err := s.handlers.IngredientSearchHandler.GetIngredientSearch(i.Ingredients)

	if err != nil && err.Error() != "sql: no rows in result set" {
		http.Error(w, "error getting ingredient search from db", http.StatusInternalServerError)
		return
	}

	if ingredientSearchInDb == nil {
		url := buildUrlWithIngredientsQuery(i.Ingredients)
		document, err := s.CallSource(url)

		if err != nil {
			http.Error(w, "bad payload", http.StatusBadRequest)
			return
		}

		recipes, err = s.app.CallRecipeResultScraping(document)

		if err != nil {
			http.Error(w, "error scraping recipe ids", http.StatusInternalServerError)
			return
		}

		newIngredientSearchId, err := s.handlers.IngredientSearchHandler.CreateIngredientSearch(i.Ingredients)

		if err != nil {
			http.Error(w, "error with db conn", http.StatusBadRequest)
			return
		}

		for _, r := range recipes {
			recipeInDb, err := s.handlers.RecipeHandler.GetRecipeById(r.Id)

			if err != nil && err.Error() != "sql: no rows in result set" {
				http.Error(w, "error getting performed query from db", http.StatusInternalServerError)
				return
			}

			if recipeInDb == nil {
				recipe := &entity.Recipe{
					Id:                 r.Id,
					Title:              r.Title,
					Description:        r.Description,
					ImageUrl:           r.ImageUrl,
					Ingredients:        r.Ingredients,
					IngredientSearchId: newIngredientSearchId,
				}
				err := s.handlers.RecipeHandler.CreateNewRecipeFromIngredients(recipe)

				if err != nil {
					http.Error(w, "error creating recipe", http.StatusInternalServerError)
					return
				}
			} else {
				err := s.handlers.RecipeHandler.UpdateRecipeIngredientSearchId(newIngredientSearchId, r.Id)
				if err != nil {
					http.Error(w, "error updating recipe", http.StatusInternalServerError)
					return
				}
			}
		}
	} else {
		recipes, err = s.handlers.RecipeHandler.GetRecipesByIngredientSearchId(int64(ingredientSearchInDb.Id))
		if err != nil {
			http.Error(w, "error with db conn", http.StatusInternalServerError)
			return
		}
	}

	j, err := json.Marshal(recipes)
	if err != nil {
		http.Error(w, "Error marshaling json response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
