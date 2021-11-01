package api

import (
	"fmt"
	"log"
	"net/http"

	domain "github.com/ionysoshedblom/go_scraper/internal/domain/interfaces"
)

type Api struct {
	api domain.ApiPort
}

func NewApi(api domain.ApiPort) *Api {
	return &Api{api: api}
}

func (a *Api) Run() {
	http.HandleFunc("/api/scraper", a.Scrape)
	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
