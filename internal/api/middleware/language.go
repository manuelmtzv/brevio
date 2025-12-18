package middleware

import (
	"net/http"

	appi18n "github.com/manuelmtzv/brevio/internal/i18n"
)

func Language(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Accept-Language")
		if lang == "" {
			lang = "en"
		}

		ctx := appi18n.WithLanguage(r.Context(), lang)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
