package api

import (
	"encoding/json"
	"fmt"
	"main/auth"
	"main/gojdb"
	"main/services"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var params map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {

		panic(err)
	}
	db := gojdb.NewGOJDB()
	encryptedbyte, _ := GetPwd(params["pass"].(string))
	params["pass"] = string(encryptedbyte)
	db.ParaClear()
	db.ParaAdd("user_id", params["user_id"])
	idexsits, _ := db.Scalar("SELECT 1 from Users where user_id = @user_id", nil)

	if idexsits != "" {
		services.ResponseWithText(w, http.StatusBadRequest, "user_id重複")
		return
	}
	rowsaffected, err := db.Insert("Users", params)
	if err != nil {
		ReturnDBError(w, err)
		return
	}
	services.ResponseWithText(w, http.StatusOK, "新增成功"+fmt.Sprint(rowsaffected))
}
func Login(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		ReturnDBError(w, err)
	}
	userID := params["user_id"]
	pass := params["pass"]
	param := make(map[string][]string)
	param["user_id"] = []string{userID}
	db := gojdb.NewGOJDB()
	row, err := db.QueryData("Select * from Users where user_id = @user_id", param)
	if err != nil {
		fmt.Println(err)
		ReturnDBError(w, err)
		return
	}
	if len(row) <= 0 {
		services.ResponseWithText(w, http.StatusUnauthorized, "user_id錯誤")
		return
	}
	usermap := row[0].(map[string]interface{})
	if usermap["user_id"].(string) == "" {
		services.ResponseWithText(w, http.StatusUnauthorized, "user_id錯誤")
	}
	if ComparePwd(usermap["pass"].(string), pass) {
		token, err := auth.SetAndGettoken(userID)
		if err != nil {
			services.ResponseWithText(w, http.StatusUnauthorized, "密碼錯誤")
		}
		setcookie(w, r, token)
		services.ResponseWithText(w, http.StatusOK, "success")
	} else {
		services.ResponseWithJson(w, http.StatusUnauthorized, "密碼錯誤")
	}

}

func Userinfo(w http.ResponseWriter, r *http.Request) {
	usercookie, _ := r.Cookie("dev-cookie")
	if usercookie == nil {
		services.ResponseWithText(w, http.StatusUnauthorized, "token empty")
		return
	}

	user, exists, _ := auth.CheckTokenExists(usercookie.Value)
	fmt.Println(user)
	if !exists {
		services.ResponseWithJson(w, http.StatusUnauthorized, "expire")
		return
	}
	bureau := user.Detail
	fmt.Println(bureau)
	services.ResponseWithJson(w, http.StatusOK, bureau)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	usercookie, _ := r.Cookie("dev-cookie")
	if usercookie == nil {
		services.ResponseWithText(w, http.StatusOK, "OK")
		return
	} else {
		manager, _ := auth.GetTokenManage()
		manager.DeleteToken(usercookie.Value)
		cookie := &http.Cookie{
			Name:     "dev-cookie",
			Value:    "",
			Expires:  time.Now().Add(-24 * time.Hour),
			HttpOnly: true,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			Path:     "/",
			Domain:   "https:192.168.2.228",
		}
		http.SetCookie(w, cookie)
		services.ResponseWithText(w, http.StatusOK, "OK")
		return
	}
}

func setcookie(w http.ResponseWriter, r *http.Request, token string) {
	cookie := &http.Cookie{
		Name:     "dev-cookie",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour), // 设置 cookie 的过期时间
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true, // 防止 JavaScript 访问 cookie
		Path:     "/",
		Domain:   "https:192.168.2.228",
	}
	http.SetCookie(w, cookie)
}

func GetPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return hash, err
}
func ComparePwd(pwd1 string, pwd2 string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}
