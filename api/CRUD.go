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

func GetRouteName(r *http.Request) map[string]string {
	vars := mux.Vars(r)
	return vars
}

func Scalar(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	gojdb := gojdb.NewGOJDB()
	sqlstring := gojdb.ScalarString(GetRouteName(r)["column"], GetRouteName(r)["table"], params)
	response, err := gojdb.Scalar(sqlstring, params)
	if err != nil {
		fmt.Println(err)
		ReturnDBError(w, err)
		return
	}
	if string(response[len(response)-4:]) == "html" {
		filedata, _ := os.ReadFile(response)
		services.ResponseWithText(w, http.StatusOK, string(filedata))
		return
	} else {
		services.ResponseWithText(w, http.StatusOK, response)
		return
	}
}
func Query(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	gojdb := gojdb.NewGOJDB()
	sqlstring := gojdb.SelectTableString(nil, GetRouteName(r)["table"], params)
	response, err := gojdb.QueryData(sqlstring, params)
	if err != nil {
		ReturnDBError(w, err)
		return
	}
	services.ResponseWithJson(w, http.StatusOK, response)
}
func Insert(w http.ResponseWriter, r *http.Request) {
	gojdb := gojdb.NewGOJDB()
	var params map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	fmt.Println(params)
	if err != nil {
		fmt.Println(err)
		services.ResponseWithText(w, http.StatusBadRequest, "malformed json data")
		return
	}

	response, err := gojdb.Insert(GetRouteName(r)["table"], params)
	if err != nil {
		ReturnDBError(w, err)
		return
	}
	services.ResponseWithJson(w, http.StatusOK, response)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	gojdb := gojdb.NewGOJDB()
	response, err := gojdb.Delete(GetRouteName(r)["table"], params)
	if err != nil {
		ReturnDBError(w, err)
		return
	}
	services.ResponseWithJson(w, http.StatusOK, response)
}

func Update(w http.ResponseWriter, r *http.Request) {

	condition := r.URL.Query()
	gojdb := gojdb.NewGOJDB()
	var params map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		fmt.Println(err)
		services.ResponseWithText(w, http.StatusBadRequest, "malformed json data")
		return
	}
	response, err := gojdb.Update(GetRouteName(r)["table"], params, condition)
	if err != nil {
		ReturnDBError(w, err)
		return
	}
	services.ResponseWithJson(w, http.StatusOK, response)
}
