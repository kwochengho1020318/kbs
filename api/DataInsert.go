package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/config"
	"main/gojdb"
	"main/services"
	"net/http"
)

func getlookup(inputstr string, Table_Name string) (map[string]interface{}, error) {
	config := config.NewConfig()
	url := config.App.ChatUrl + "api/chat/check?Table_Name=" + Table_Name
	payload := []byte(inputstr)
	resp, err := http.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var lookup map[string]interface{}
	result, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(result, &lookup)
	if err != nil {

		return nil, err
	}
	return lookup, nil
}

// 給定json string 及table ,新增資料
func datainsert(w http.ResponseWriter, table string, content []byte, lookup map[string]interface{}) error {
	db := gojdb.NewGOJDB()
	var isObject bool
	if table == "" {
		services.ResponseWithText(w, 400, "Table_Name not provided")
	}
	//unescaped, _ := strconv.Unquote(string(content))
	fmt.Println(string(content))
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
			return err
		}
		log := ""

		recursiveParse(params, table, db, &log, lookup)

		services.ResponseWithText(w, 200, log)
		return err
	} else {
		var params []interface{}

		err := json.Unmarshal(content, &params)
		if err != nil {
			fmt.Println(string(content))
			fmt.Println(err)
			services.ResponseWithText(w, http.StatusBadRequest, err.Error())
			return err
		}
		log := ""
		for _, v := range params {

			recursiveParse(v.(map[string]interface{}), table, db, &log, lookup)
		}

		services.ResponseWithText(w, 200, log)
	}
	return nil
}

func recursiveParse(json map[string]interface{}, table string, db *gojdb.GOJDB, log *string, lookup map[string]interface{}) {
	attributes := make(map[string]interface{})
	for key, value := range json {
		if _, ok := value.([]interface{}); ok {
			for _, entry := range value.([]interface{}) {
				recursiveParse(entry.(map[string]interface{}), table, db, log, lookup)
			}
		} else if _, ok := value.(map[string]interface{}); ok {
			recursiveParse(value.(map[string]interface{}), table, db, log, lookup)
		} else {
			alteredkey := lookup[key]
			if alteredkey != nil {

				attributes[alteredkey.(string)] = value
			} else {
				attributes[key] = value
			}
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
	} else {
		fmt.Println("insert into " + table)
	}
}
