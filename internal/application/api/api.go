package api

import (
	"io"
)

type Application struct {
	scraper ScraperPort
}

func NewApplication(scraper ScraperPort) *Application {
	return &Application{ scraper: scraper }
}

func (a Application) HandleSource(src io.ReadCloser) (string, error) {
	stringSrc, err := a.scraper.HandleSource(src)

	if err != nil {
		return "", err
	}

	return stringSrc, nil
}
