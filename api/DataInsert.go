package api

import (
	"encoding/json"
	"fmt"
	"io"
	"main/gojdb"
	"main/services"
	"net/http"
)

func RssInsert(w http.ResponseWriter, r *http.Request) {
	db := gojdb.NewGOJDB()

	content, _ := io.ReadAll(r.Body)
	var isObject bool
	if string((content)[0]) == "{" {
		isObject = true
	} else {
		isObject = false
	}
	if isObject {
		var params map[string]interface{}

		err := json.Unmarshal(content, &params)
		if err != nil {
			//fmt.Println(string(content))
			fmt.Println(err)
			services.ResponseWithText(w, http.StatusBadRequest, "malformed json data")
			return
		}
		log := ""

		recursiveParse(params, "", db, &log)

		services.ResponseWithText(w, 200, log)
	} else {
		var params []interface{}

		err := json.Unmarshal(content, &params)
		if err != nil {
			fmt.Println(string(content))
			fmt.Println(err)
			services.ResponseWithText(w, http.StatusBadRequest, err.Error())
			return
		}
		log := ""
		for _, v := range params {

			recursiveParse(v.(map[string]interface{}), "", db, &log)
		}

		services.ResponseWithText(w, 200, log)
	}

}
func recursiveParse(json map[string]interface{}, table string, db *gojdb.GOJDB, log *string) {
	attributes := make(map[string]interface{})
	for key, value := range json {
		if _, ok := value.([]interface{}); ok {
			for _, entry := range value.([]interface{}) {
				recursiveParse(entry.(map[string]interface{}), key, db, log)
			}
		} else if _, ok := value.(map[string]interface{}); ok {
			recursiveParse(value.(map[string]interface{}), key, db, log)
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
		*log += "table not found ,skip inserting " + table + "\n"
	}
	fmt.Println("insert into " + table)
}
