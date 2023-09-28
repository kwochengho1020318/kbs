package api

import (
	"encoding/json"
	"fmt"
	"main/gojdb"
	"main/services"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func GetRender(w http.ResponseWriter, r *http.Request) {

	//usercookie, _ := r.Cookie("dev-cookie")
	// if usercookie == nil {
	// 	//services.ResponseWithText(w, http.StatusUnauthorized, "no cookie")
	// 	//return

	// }
	// _, _, err := auth.CheckTokenExists(usercookie.Value)
	// if err != nil {
	// 	fmt.Println(err)
	// 	ReturnDBError(w, err)
	// 	return
	// }
	if true {
		page := mux.Vars(r)["page"]

		pagefile, err := os.ReadFile("files/pages/" + page + ".4u")
		if err != nil {
			fmt.Println(err)
			services.ResponseWithJson(w, http.StatusBadRequest, ErrorCode{101, "page not found "})
			return
		}
		services.ResponseWithHtml(w, http.StatusOK, pagefile)
	} else {

		services.ResponseWithHtml(w, http.StatusOK, []byte("token expired"))
		return
	}
}
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

	usercookie, _ := r.Cookie("dev-cookie")
	if usercookie == nil {
		filestring, err := ExpirePage()
		if err != nil {
			services.ResponseWithHtml(w, http.StatusInternalServerError, []byte(""))
		}
		services.ResponseWithHtml(w, http.StatusOK, []byte(filestring))
		return
	}
	db := gojdb.NewGOJDB()
	db.ParaClear()
	db.ParaAdd("token", usercookie.Value)
	existsStr := "select user_id from [token] where token = @token"
	userinfo, err := db.QueryJsonWithParameters(existsStr)
	if err != nil {
		fmt.Println(err)
	}
	if len(userinfo) > 0 {
		page := mux.Vars(r)["page"]

		pagefile, err := os.ReadFile("templatesite/" + page + ".html")
		Person, _ := json.Marshal(userinfo[0])
		pagestring := string(pagefile)
		pagestring = fmt.Sprintf("<script>var Person = %s </script>,", Person) + pagestring

		if err != nil {
			fmt.Println(err)
			services.ResponseWithJson(w, http.StatusBadRequest, ErrorCode{101, "page not found "})
			return
		}
		services.ResponseWithHtml(w, http.StatusOK, []byte(pagestring))
	} else {

		filestring, err := ExpirePage()
		if err != nil {
			services.ResponseWithHtml(w, http.StatusInternalServerError, []byte(""))
		}
		services.ResponseWithHtml(w, http.StatusOK, []byte(filestring))
		return
	}
}
