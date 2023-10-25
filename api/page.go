package api

import (
	"fmt"
	"main/services"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func PageGetter(w http.ResponseWriter, r *http.Request) {

	page := mux.Vars(r)["page"]
	if page == "" {
		page = "index"
	}
	pagefile, err := os.ReadFile("templatesite/" + page + ".html")
	pagestring := string(pagefile)

	if err != nil {
		fmt.Println(err)
		services.ResponseWithJson(w, http.StatusBadRequest, ErrorCode{101, "page not found "})
		return
	}
	services.ResponseWithHtml(w, http.StatusOK, []byte(pagestring))

}
