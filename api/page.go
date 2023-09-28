package api

import (
	"encoding/json"
	"fmt"
	"main/config"
	"main/gojdb"
	"main/services"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func ExpirePage() (string, error) {
	pagefile, err := os.ReadFile("templatesite/index.html")
	if err != nil {
		return "", err
	}
	filestring := string(pagefile)
	filestring = "<script>alert('token expire')</script>" + filestring
	return filestring, nil
}

func PageGetter(w http.ResponseWriter, r *http.Request) {
	config := config.NewConfig("appsettings.json")
	usercookie, _ := r.Cookie("Token")
	db := gojdb.NewGOJDB()
	db.ParaClear()
	db.ParaAdd("token", usercookie.Value)
	existsStr := "select user_id from [token] where token = @token"
	userinfo, err := db.QueryJsonWithParameters(existsStr)
	if err != nil {
		ReturnDBError(w, err)
		return
	}
	if len(userinfo) > 0 {
		page := mux.Vars(r)["page"]
		pagefile, err := os.ReadFile("templatesite/" + page + ".html")
		Person, _ := json.Marshal(userinfo[0])
		pagestring := string(pagefile)
		pagestring = fmt.Sprintf("<script>var User = %s </script>,", Person) + pagestring

		if err != nil {
			fmt.Println(err)
			services.ResponseWithJson(w, http.StatusBadRequest, ErrorCode{101, "page not found "})
			return
		}
		services.ResponseWithHtml(w, http.StatusOK, []byte(pagestring))
		return
	} else {
		http.Redirect(w, r, config.App.LoginSite, http.StatusSeeOther)
		return
	}
}
