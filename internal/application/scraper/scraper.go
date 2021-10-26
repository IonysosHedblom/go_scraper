package scraper

import (
	"io"
	"io/ioutil"
)

type Scraper struct {}

func New() *Scraper {
	return &Scraper{}
}

func (s Scraper) HandleSource(src io.ReadCloser) (string, error) {
	dataInBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return "", err
	}

	defer src.Close()
	htmlStr := string(dataInBytes)
	return htmlStr, nil
}
