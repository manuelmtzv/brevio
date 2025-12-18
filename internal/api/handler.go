package api

import "net/http"

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func Adapt(h HandlerFunc, onError ErrorHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			onError(w, r, err)
		}
	}
}
