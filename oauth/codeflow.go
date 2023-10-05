package oauth

import (
	"main/config"
	"net/http"
)

func OauthStart(w http.ResponseWriter, r *http.Request) {
	config := config.NewConfig("appsettings.json")
	http.Redirect(w, r, config.App.LoginSite, http.StatusSeeOther)
}
