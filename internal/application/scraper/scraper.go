package scraper

import (
	"bytes"

	"golang.org/x/net/html"
)

type Scraper struct{}

func New() *Scraper {
	return &Scraper{}
}

func (s Scraper) HandleSource(src *html.Node) ([]*bytes.Buffer, error) {
	var results []*bytes.Buffer
	var visitNode func(*html.Node)
	visitNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Parent.Data == "h2" {
			text := &bytes.Buffer{}
			s.CollectText(n, text)

			results = append(results, text)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}
	s.ForEachNode(src, visitNode, nil)
	return results, nil
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
