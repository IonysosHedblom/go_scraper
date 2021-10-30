package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func (s Server) Scrape(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return
	}

	qs := req.URL.Query()
	if len(qs) > 1 {
		http.Error(w, "too many queries", http.StatusBadRequest)
		return
	}

	q := qs["query"]
	if len(q) > 1 {
		http.Error(w, "too much queries", http.StatusBadRequest)
		return
	}

	url := buildUrl(q[0])

	document, err := s.CallSource(url)

	if err != nil {
		http.Error(w, "bad payload", http.StatusBadRequest)
		return
	}

	response, err := s.api.HandleSource(document)

	if err != nil {
		http.Error(w, "something went wrong in api layer", http.StatusBadRequest)
		return
	}

	j, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func (s Server) CallSource(url string) (*html.Node, error) {
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

func buildUrl(query string) string {
	url := fmt.Sprintf("https://www.ica.se/Templates/ajaxresponse.aspx?ajaxFunction=RecipeListMdsa&mdsarowentityid=&num=16&query=%s&sortbymetadata=Relevance&id=12", query)
	return url
}
