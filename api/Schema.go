package api

import (
	"encoding/json"
	"main/gojdb"
	"main/services"
	"net/http"
)

func UpdateTable(w http.ResponseWriter, r *http.Request) {
	gojdb := gojdb.NewGOJDB()
	var params map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		services.ResponseWithText(w, http.StatusBadRequest, "malformed json data")
		return
	}
	gojdb.UpdateTable(params)

}
func UpdateView(w http.ResponseWriter, r *http.Request) {
	gojdb := gojdb.NewGOJDB()
	var params map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		services.ResponseWithText(w, http.StatusBadRequest, "malformed json data")
		return
	}
	gojdb.UpdateView(params)
}
