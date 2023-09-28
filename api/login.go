package api

import (
	"net/http"
	"time"
)

func setcookie(w http.ResponseWriter, r *http.Request, token string) {
	cookie := &http.Cookie{
		Name:     "dev-cookie",
		Value:    "B6A95CA92D477CAFBD5DEABC76ED82387E65E23424F1C946966987A76AF78054",
		Expires:  time.Now().Add(24 * time.Hour), // 设置 cookie 的过期时间
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true, // 防止 JavaScript 访问 cookie
		Path:     "/",
	}
	http.SetCookie(w, cookie)
}
