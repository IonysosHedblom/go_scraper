package service

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ScraperService interface {
	ConvertToString(io.ReadCloser) (string)
	FindElement(content string, startTag string, endTag string) ([]byte)
}

type scraperService struct {}

func NewScraperService () ScraperService {
	return &scraperService{}
}

func (w *scraperService) ConvertToString(body io.ReadCloser) string  {
	dataInBytes, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}

	defer body.Close()
	htmlStr := string(dataInBytes)
	fmt.Println(htmlStr)
	return htmlStr
}

func (w *scraperService) FindElement(content string, startTag string, endTag string) ([]byte) {
	elStart := strings.Index(content, startTag)
	noElStart := elStart == -1
	if noElStart {
		fmt.Printf("No %v startTag found", startTag)
		os.Exit(0)
	}

	elEnd := strings.Index(content, endTag)
	noElEnd := elEnd == -1
	
	if noElEnd {
		fmt.Printf("No %v endTag found", endTag)
		os.Exit(0)
	}

	pageTag := []byte(content[elStart:elEnd])
	return pageTag
}
