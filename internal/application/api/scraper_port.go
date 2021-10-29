package api

import (
	"bytes"

	"golang.org/x/net/html"
)

type ScraperPort interface {
	HandleSource(src *html.Node) ([]*bytes.Buffer, error)
}
