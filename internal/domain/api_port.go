package domain

type ApiPort interface {
	HandleSource(src []string) (string, error)
}
