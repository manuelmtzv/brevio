package request

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Param(r *http.Request, key string) (string, error) {
	v := chi.URLParam(r, key)
	if v == "" {
		return "", ErrMissingParam{Key: key}
	}
	return v, nil
}

func QueryInt(r *http.Request, key string, def int) (int, error) {
	v := r.URL.Query().Get(key)
	if v == "" {
		return def, nil
	}
	return strconv.Atoi(v)
}
