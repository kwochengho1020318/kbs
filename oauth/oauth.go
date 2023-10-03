package oauth

import (
	"encoding/json"
	"main/api"
	"main/gojdb"
	"net/http"
)

func CallBack(w http.ResponseWriter, r *http.Request) {
	var params map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		api.ReturnDBError(w, err)
		return
	}
	if params["code"].(string) == "200" {
		db := gojdb.NewGOJDB()
		redirectUrl := r.URL.Query().Get("redirectUrl")
		delete(params, "code")
		_, err := db.Insert("Token", params)
		if err != nil {
			api.ReturnDBError(w, err)
			return
		}
		http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
	} else {
		return
	}

}
