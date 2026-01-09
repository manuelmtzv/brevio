package presenters

import "fmt"

type HTTPURLBuilder struct {
	BaseURL string
}

func NewHTTPURLBuilder(baseURL string) *HTTPURLBuilder {
	return &HTTPURLBuilder{BaseURL: baseURL}
}

func (b *HTTPURLBuilder) ShortURL(code string) string {
	return fmt.Sprintf("%s/%s", b.BaseURL, code)
}
