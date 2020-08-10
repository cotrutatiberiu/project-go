package middleware

import (
	"net/http"
)

func adminOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if true {
			http.NotFound(w, r)
			return
		}
		h(w, r)
	}
}
