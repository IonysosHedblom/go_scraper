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

	performedQueryInDb, err := s.app.GetPerformedQuery(q[0])

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

		newQueryId, err := s.app.CreateNewPerformedQuery(q[0])
		if err != nil {
			http.Error(w, "error with db conn", http.StatusBadRequest)
			return
		}

		for _, r := range recipes {
			recipeInDb, err := s.app.GetRecipeById(r.Id)

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
				err := s.app.CreateNewRecipe(recipe)

				if err != nil {
					http.Error(w, "error creating new recipe", http.StatusInternalServerError)
					return
				}
			} else {
				err := s.app.UpdateRecipeQueryId(newQueryId, r.Id)
				if err != nil {
					http.Error(w, "error updating recipe", http.StatusInternalServerError)
					return
				}
			}
		}
	} else {
		recipes, err = s.app.GetRecipesByQueryId(int64(performedQueryInDb.Id))
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

func (s *api) PostWithIngredients(w http.ResponseWriter, req *http.Request) {
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

	ingredientSearchInDb, err := s.app.GetIngredientSearch(i.Ingredients)

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

		newIngredientSearchId, err := s.app.CreateIngredientSearch(i.Ingredients)

		if err != nil {
			http.Error(w, "error with db conn", http.StatusBadRequest)
			return
		}

		for _, r := range recipes {
			recipe := &entity.Recipe{
				Id:                 r.Id,
				Title:              r.Title,
				Description:        r.Description,
				ImageUrl:           r.ImageUrl,
				Ingredients:        r.Ingredients,
				IngredientSearchId: newIngredientSearchId,
			}
			err = s.app.CreateNewRecipeFromIngredients(recipe)
		}

		if err != nil {
			http.Error(w, "error with db conn", http.StatusInternalServerError)
			return
		}
	} else {
		recipes, err = s.app.GetRecipesByIngredientSearchId(int64(ingredientSearchInDb.Id))
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
	url := fmt.Sprintf("https://www.ica.se/Templates/ajaxresponse.aspx?ajaxFunction=RecipeListMdsa&mdsarowentityid=&num=16&query=%s&sortbymetadata=Relevance&id=12", query)
	return url
}
