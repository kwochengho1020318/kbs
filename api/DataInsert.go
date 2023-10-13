package api

import (
	"encoding/json"
	"fmt"
	"main/gojdb"
	"main/services"
	"net/http"
)

func RssInsert(w http.ResponseWriter, r *http.Request) {
	db := gojdb.NewGOJDB()
	var params map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		fmt.Println(err)
		services.ResponseWithText(w, http.StatusBadRequest, "malformed json data")
		return
	}
	recursiveParse(params, "", db)
}
func recursiveParse(json map[string]interface{}, table string, db *gojdb.GOJDB) {
	attributes := make(map[string]interface{})
	for key, value := range json {
		if _, ok := value.([]interface{}); ok {
			for _, entry := range value.([]interface{}) {
				recursiveParse(entry.(map[string]interface{}), key, db)
			}
		} else if _, ok := value.(map[string]interface{}); ok {
			recursiveParse(value.(map[string]interface{}), key, db)
		} else {
			attributes[key] = value
		}
	}
	//fmt.Println(attributes)
	if table == "" {
		return
	}
	_, err := db.Insert(table, attributes)
	if err != nil {
		fmt.Println(err)
		fmt.Println("table not found ,skip inserting " + table)

	}
	return
}
