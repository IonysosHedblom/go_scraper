package api

import "net/http"

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

}
