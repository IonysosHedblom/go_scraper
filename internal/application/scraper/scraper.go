package scraper

import (
	"bytes"
	"fmt"

	"github.com/ionysoshedblom/go_scraper/internal/domain/helpers"
	"golang.org/x/net/html"
)

type Scraper struct{}

func New() *Scraper {
	return &Scraper{}
}

func (s Scraper) HandleSource(src *html.Node) ([]byte, error) {
	var results []string
	var visitNode func(*html.Node)
	visitNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Parent.Data == "h2" {
			text := &bytes.Buffer{}
			s.CollectText(n, text)
			fmt.Println(text)
			results = append(results, n.Data)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}
	s.ForEachNode(src, visitNode, nil)
	bytes := helpers.StringSliceToByteSlice(results)
	return bytes, nil
}

func (s Scraper) ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s.ForEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func (s Scraper) CollectText(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s.CollectText(c, buf)
	}
}
