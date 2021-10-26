package scraper

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
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

func (s Scraper) GetElement(src string, sTag string, eTag string) ([]byte, error) {
	elStart := strings.Index(src, sTag)
	noElStart := elStart == -1
	if noElStart {
		return nil, fmt.Errorf("no %v starttag found", sTag)
	}

	elEnd := strings.Index(src, eTag)
	noElEnd := elEnd == -1
	
	if noElEnd {
		return nil, fmt.Errorf("no %v endtag found", eTag)
	}

	pageTag := []byte(src[elStart:elEnd])
	return pageTag, nil
}
