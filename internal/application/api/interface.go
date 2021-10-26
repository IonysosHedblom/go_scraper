package api

import "io"

type Scraper interface {
	GetSource(url string) (io.ReadCloser, error)
	HandleSource(io.ReadCloser) (string, error)
	ExtractElements(startTag string, endTag string) (string, error)
}