package domain

import "io"


type ApiPort interface {
	GetSource(url string) (io.ReadCloser, error)
	HandleSource(src io.ReadCloser) (string, error)
}
