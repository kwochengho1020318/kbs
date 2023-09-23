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
	err = gojdb.UpdateTable(params)
	if err != nil {
		ReturnDBError(w, err)
	}

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
func UpdateStoredProcedure(w http.ResponseWriter, r *http.Request) {
	gojdb := gojdb.NewGOJDB()
	var params map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		services.ResponseWithText(w, http.StatusBadRequest, "malformed json data")
		return
	}
	gojdb.UpdateStoreProcedure(params)
}
func UpdateSchema(w http.ResponseWriter, r *http.Request) {
	gojdb := gojdb.NewGOJDB()
	var params map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		services.ResponseWithText(w, http.StatusBadRequest, "malformed json data")
		return
	}
	Tables := params["Tables"].([]interface{})
	for _, table := range Tables {
		gojdb.UpdateTable(table.(map[string]interface{}))
	}
	Views := params["Views"].([]interface{})
	for _, view := range Views {
		gojdb.UpdateView(view.(map[string]interface{}))
	}
	Procedures := params["Procedures"].([]interface{})
	for _, proc := range Procedures {
		gojdb.UpdateStoreProcedure(proc.(map[string]interface{}))
	}
	services.ResponseWithText(w, http.StatusOK, "success")

}
