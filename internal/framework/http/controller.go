package server

import (
	"io/ioutil"
	"net/http"
)

func (s Server) Scrape(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(w, "cant read body", http.StatusBadRequest)
		return
	}

	response, err := s.CallSource(string(body))

	if err != nil {
		http.Error(w, "bad payload", http.StatusBadRequest)
		return
	}
	
	stringRes, err := s.api.HandleSource(response)
	
	if err != nil {
		http.Error(w,"something went wrong in api layer", http.StatusBadRequest)
		return
	}

	w.Write([]byte(stringRes))
}

// func (s Server) CallSource(url string) ([]string, error) {
// 	res, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("getting %s: %s", url, res.Status)
// 	}

// 	doc, err := html.Parse(res.Body)

// 	if err != nil {
// 		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
// 	}

// 	return nil, nil
// }

