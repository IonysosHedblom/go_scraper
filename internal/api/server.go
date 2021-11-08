package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ionysoshedblom/go_scraper/internal/domain/abstractions"
)

type api struct {
	app abstractions.AppPort
}

func NewApi(app abstractions.AppPort) *api {
	return &api{app: app}
}

func (a *api) Run() {
	http.HandleFunc("/api/scraper", a.ScraperRouter)
	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
