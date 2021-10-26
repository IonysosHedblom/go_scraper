package api

import (
	"io"
	"log"
	"net/http"
)

type Application struct {
	scraper Scraper
}

func NewApplication(scraper Scraper) *Application {
	return &Application{}
}

func (a Application) GetSource(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return response.Body, nil
}

func (a Application) HandleSource(src io.ReadCloser) (string, error) {
	stringSrc, err := a.scraper.HandleSource(src)

	if err != nil {
		return "", err
	}

	return stringSrc, nil
}
