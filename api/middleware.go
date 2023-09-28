package api

import (
	"main/config"
	"net/http"
)

func CorsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://192.168.2.228")
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		w.Header().Set("Access-Control-Allow-Origin", "null")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
func AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config := config.NewConfig("appsettings.json")

		usercookie, _ := r.Cookie("Token")
		if usercookie == nil {

			http.Redirect(w, r, config.App.LoginSite, http.StatusSeeOther)
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
