package domain

import (
	"bytes"

	"golang.org/x/net/html"
)

type ApiPort interface {
	HandleSource(*html.Node) ([]*bytes.Buffer, error)
}
