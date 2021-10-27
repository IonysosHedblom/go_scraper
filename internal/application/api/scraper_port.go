package api

import "io"

type ScraperPort interface {
	HandleSource(io.ReadCloser) (string, error)
	GetElement(src string, sTag string, eTag string) ([]byte, error)
	ExtractElements(startTag string, endTag string) (string, error)
}