package api

import "golang.org/x/net/html"

type ScraperPort interface {
	HandleSource(src *html.Node) ([]byte, error)
	ForEachNode(n *html.Node, pre, post func(n *html.Node))
}