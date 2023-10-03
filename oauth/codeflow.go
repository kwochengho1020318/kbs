package oauth

import "net/http"

func OauthStart(w http.ResponseWriter, r *http.Request) {
	
	http.Redirect(w, r, "https://loginsite.com", http.StatusSeeOther)
}
