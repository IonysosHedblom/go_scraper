package scraper

import (
	"fmt"

	"golang.org/x/net/html"
)

type Scraper struct {}

func New() *Scraper {
	return &Scraper{}
}

func (s Scraper) HandleSource(src *html.Node) ([]string, error) {	
	var results []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "h2" {
			for _, a := range n.Attr {
				fmt.Println(a.Val)
				results = append(results, a.Val)
			}
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
