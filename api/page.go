package api

import (
	"main/auth"
	"main/services"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func PageGetter(w http.ResponseWriter, r *http.Request) {
	usercookie, _ := r.Cookie("dev-cookie")
	if usercookie == nil {
		services.ResponseWithText(w, http.StatusUnauthorized, "no cookie")

	} else {
		_, exists, err := auth.CheckTokenExists(usercookie.Value)
		if err != nil {
			return
		}
		if exists {
			page := mux.Vars(r)["page"]
			// a := gojdb.NewGOJDB()
			// slice := make(map[string][]string)
			// slice["filename"] = []string{page}
			// ID, err := a.Scalar("Select user_id from pages where filename = @filename", slice)
			// if ID != user.Id {
			// 	services.ResponseWithText(w, http.StatusUnauthorized, "Not your Page")
			// 	return
			// }
			// if err != nil {
			// 	panic(err)
			// }

			pagefile, err := os.ReadFile("files/pages/" + page + ".4u")
			if err != nil {
				services.ResponseWithJson(w, http.StatusBadRequest, ErrorCode{101, "page not found "})
				return
			}
			services.ResponseWithHtml(w, http.StatusOK, pagefile)
		} else {
			services.ResponseWithText(w, http.StatusUnauthorized, "token expire")
			return
		}
	}

}
