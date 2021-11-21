package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
)

func (s *api) InventoryRouter(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		s.GetItems(w, req)
	case "POST":
		s.AddItems(w, req)
	default:
		http.Error(w, "No support for this method", http.StatusMethodNotAllowed)
		return
	}
}

func (s *api) GetItems(w http.ResponseWriter, req *http.Request) {
	queries := req.URL.Query()
	if len(queries) != 1 {
		http.Error(w, "bad query param", http.StatusBadRequest)
		return
	}

	q := queries["user"]
	if len(q) != 1 {
		http.Error(w, "bad query param", http.StatusBadRequest)
		return
	}

	itemsFromDb, err := s.handlers.InventoryHandler.Get(q[0])

	if err != nil {
		http.Error(w, "Error fetching items from db for this user", http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(itemsFromDb)

	if err != nil {
		http.Error(w, "Error with marshaling json from InventoryRouter", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func (s *api) AddItems(w http.ResponseWriter, req *http.Request) {
	var reqBody entity.AddItemsRequest

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

	json.Unmarshal(body, &reqBody)

	err = s.handlers.InventoryHandler.AddItems(reqBody.Items)

	if err != nil {
		http.Error(w, "Error with adding items to database from inventory router", http.StatusBadRequest)
		return
	}

	j, err := json.Marshal(reqBody)

	if err != nil {
		http.Error(w, "Error marshaling json response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)

}
