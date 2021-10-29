package scraper

import (
	"bytes"

	"github.com/ionysoshedblom/go_scraper/internal/domain/entity"
	"golang.org/x/net/html"
)

type Scraper struct{}

func New() *Scraper {
	return &Scraper{}
}

func (s Scraper) HandleSource(n *html.Node) ([]entity.Recipe, error) {
	var b []*bytes.Buffer
	var visitNode func(*html.Node)
	visitNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Parent.Data == "h2" {
			text := &bytes.Buffer{}
			writeNodeContentToBuffer(n, text)
			b = append(b, text)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}

	forEachNode(n, visitNode, nil)
	recipeTitles := mapBufValuesToStruct(b)

	return recipeTitles, nil
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

func mapBufValuesToStruct(barr []*bytes.Buffer) []entity.Recipe {
	var out []entity.Recipe
	for _, buf := range barr {
		recipe := &entity.Recipe{Title: buf.String()}
		out = append(out, *recipe)
	}
	return out
}
