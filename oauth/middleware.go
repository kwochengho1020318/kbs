package oauth

import (
	"main/gojdb"
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

		usercookie, _ := r.Cookie("Token")
		if usercookie == nil {
			OauthStart(w, r)
			return
		} else {
			db := gojdb.NewGOJDB()
			db.ParaAdd("Token", usercookie.Value)
			v, _ := db.Scalar("Select User_ID from Token where Token = @Token", nil)
			if v == "" {
				OauthStart(w, r)
				return
			} else {
				next.ServeHTTP(w, r)
				return
			}
		}
	})
}
