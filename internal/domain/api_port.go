package domain

import "golang.org/x/net/html"

type ApiPort interface {
	HandleSource(*html.Node) ([]byte, error)
}
