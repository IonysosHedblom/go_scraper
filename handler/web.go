package handler

import (
	"io"
	"log"
	"net/http"
)


func (m *MicroService) Get(url string) io.ReadCloser {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return response.Body
}