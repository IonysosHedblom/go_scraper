package domain

type ApiPort interface {
	HandleSource([]string) (string, error)
}
