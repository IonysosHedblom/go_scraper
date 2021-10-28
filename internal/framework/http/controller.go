package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

type urlRequestBody struct {
	Url string
}

func (s Server) Scrape(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "error", http.StatusBadRequest)
	}
	var requestBody urlRequestBody

	json.Unmarshal(body ,&requestBody)

	if err != nil {
		http.Error(w, "error marshaling body", http.StatusBadRequest)
	}

	response, err := s.CallSource(requestBody.Url)

	if err != nil {
		http.Error(w, "bad payload", http.StatusBadRequest)
		return
	}
	
	html, err := s.api.HandleSource(response)
	
	if err != nil {
		http.Error(w,"something went wrong in api layer", http.StatusBadRequest)
		return
	}

	w.Write([]byte(html))
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

	defer res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	
	return doc, nil
}

