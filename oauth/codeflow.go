package oauth

import (
	"fmt"
	"main/config"
	"net/http"
)

func OauthStart(w http.ResponseWriter, r *http.Request) {
	config := config.NewConfig()
	http.Redirect(w, r, config.App.LoginSite+fmt.Sprintf("?user_p=%s_%s", config.App.User, config.App.Pid), http.StatusSeeOther)
}
