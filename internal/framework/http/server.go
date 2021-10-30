package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ionysoshedblom/go_scraper/internal/domain"
)

type Server struct {
	api domain.ApiService
}

func NewServer(api domain.ApiService) *Server {
	return &Server{api: api}
}

func (httpServer *Server) Run() {
	http.HandleFunc("/api/scraper", httpServer.Scrape)
	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
