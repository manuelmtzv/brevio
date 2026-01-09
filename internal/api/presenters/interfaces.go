package presenters

type URLBuilder interface {
	ShortURL(code string) string
}
