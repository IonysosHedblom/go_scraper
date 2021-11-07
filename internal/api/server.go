package api

import (
	"fmt"
	"log"
	"net/http"

	domain "github.com/ionysoshedblom/go_scraper/internal/domain/interfaces"
)

type api struct {
	api domain.ApiPort
}

func NewApi(a domain.ApiPort) *api {
	return &api{api: a}
}

func (a *api) Run() {
	http.HandleFunc("/api/scraper", a.GetQuery)
	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
