package scraper

import (
	"bytes"
	"fmt"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type Scraper struct{}

func New() *Scraper {
	return &Scraper{}
}

func (s Scraper) HandleSource(n *html.Node) ([]entity.Recipe, error) {
	var titles []*bytes.Buffer
	var desc []*bytes.Buffer
	var visitNode func(*html.Node)

	visitNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Parent.Data == "h2" {
			titleBuf := &bytes.Buffer{}
			writeNodeContentToBuffer(n, titleBuf)
			titles = append(titles, titleBuf)
		}

		if n.Type == html.ElementNode && n.Parent.Data == "a" && n.Data == "p" {
			dBuf := &bytes.Buffer{}
			writeNodeContentToBuffer(n, dBuf)
			desc = append(desc, dBuf)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}

	forEachNode(n, visitNode, nil)
	recipes := mapBufValuesToStruct(titles, desc)
	fmt.Println(recipes)
	return recipes, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func writeNodeContentToBuffer(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		writeNodeContentToBuffer(c, buf)
	}
}

func mapBufValuesToStruct(titles []*bytes.Buffer, descriptions []*bytes.Buffer) []entity.Recipe {
	var out []entity.Recipe
	for i := 0; i < len(titles); i++ {
		recipe := &entity.Recipe{Title: titles[i].String(), Description: descriptions[i].String()}
		out = append(out, *recipe)
	}
	return out
}
