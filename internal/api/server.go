package api

import (
	"fmt"
	"log"
	"net/http"

	domain "github.com/ionysoshedblom/go_scraper/internal/domain/interfaces"
)

type Server struct {
	api domain.ApiPort
}

func NewServer(api domain.ApiPort) *Server {
	return &Server{api: api}
}

func (httpServer *Server) Run() {
	http.HandleFunc("/api/scraper", httpServer.Scrape)
	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
