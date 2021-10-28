package api

import "golang.org/x/net/html"

type ScraperPort interface {
	HandleSource(*html.Node) (string, error)
	GetElement(string, string, string) ([]byte, error)
	ExtractElements(string, string) (string, error)
}