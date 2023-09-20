package api

import (
	"net/http"
)

func CorsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://192.168.2.228")
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
