package server

import (
	"fmt"
	"net/http"

	"github.com/ionysoshedblom/go_scraper/internal/domain"
)

type Server struct {
	api domain.ApiPort
}

func NewServer(api domain.ApiPort) *Server {
	return &Server{ api: api }
}

func (httpServer *Server) Run() {
	http.HandleFunc("/api/scraper", httpServer.Scrape)
	fmt.Println("Server running on port 1323")
	http.ListenAndServe(":1323", nil)
}
