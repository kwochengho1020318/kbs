package api

import (
	"fmt"
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

		services.ResponseWithText(w, http.StatusUnauthorized, "token expire")
		return
	}
}
func PageGetter(w http.ResponseWriter, r *http.Request) {

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

		pagefile, err := os.ReadFile("templatesite/" + page)
		if err != nil {
			fmt.Println(err)
			services.ResponseWithJson(w, http.StatusBadRequest, ErrorCode{101, "page not found "})
			return
		}
		services.ResponseWithHtml(w, http.StatusOK, pagefile)
	} else {

		services.ResponseWithText(w, http.StatusUnauthorized, "token expire")
		return
	}
}

func Src(w http.ResponseWriter, r *http.Request) {

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
		srctype := mux.Vars(r)["srctype"]
		srcfile := mux.Vars(r)["srcfile"]
		filepath := fmt.Sprintf("templatesite/assets/%s/%s", srctype, srcfile)
		fmt.Println(filepath)
		pagefile, err := os.ReadFile(filepath)
		if err != nil {
			fmt.Println(err)
			services.ResponseWithJson(w, http.StatusBadRequest, ErrorCode{101, "page not found "})
			return
		}
		services.ResponseWithHtml(w, http.StatusOK, pagefile)
	} else {

		services.ResponseWithText(w, http.StatusUnauthorized, "token expire")
		return
	}

}
