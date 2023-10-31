package source

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"main/config"
	"main/gojdb"
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

func datainsert(w http.ResponseWriter, table string, content []byte, lookup map[string]interface{}) (string, error) {
	db := gojdb.NewGOJDB()
	var isObject bool
	if table == "" {
		return "", errors.New("Table_Name not provided ")
	}
	//unescaped, _ := strconv.Unquote(string(content))
	if string((content)[0]) == "{" {
		isObject = true
	} else {
		isObject = false
	}
	log := ""
	var rowsaffected int
	if isObject {
		var params map[string]interface{}

		err := json.Unmarshal(content, &params)
		if err != nil {
			//fmt.Println(string(content))
			log += "malformed json\n"
			return log, err
		}

		recursiveParse(params, table, db, &log, &rowsaffected, lookup)

		log += "Rows Affected : " + fmt.Sprint(rowsaffected) + "\n"
		return log, err
	} else {
		var params []interface{}

		err := json.Unmarshal(content, &params)
		if err != nil {
			log += err.Error() + "\n"
			return log, err
		}

		for _, v := range params {

			recursiveParse(v.(map[string]interface{}), table, db, &log, &rowsaffected, lookup)
		}

		log += "Rows Affected : " + fmt.Sprint(rowsaffected) + "\n"
	}
	//fmt.Println(log)
	return log, nil
}

func recursiveParse(json map[string]interface{}, table string, db *gojdb.GOJDB, log *string, rowsaffected *int, lookup map[string]interface{}) {
	attributes := make(map[string]interface{})
	db.ParaClear()
	for key, value := range json {
		if _, ok := value.([]interface{}); ok {
			for _, entry := range value.([]interface{}) {
				recursiveParse(entry.(map[string]interface{}), table, db, log, rowsaffected, lookup)
			}
		} else if _, ok := value.(map[string]interface{}); ok {
			recursiveParse(value.(map[string]interface{}), table, db, log, rowsaffected, lookup)
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
		*log += "warning : one insert failed" + "\n"
	} else {
		*rowsaffected += 1
		succdssinfo := "insert success\n"
		*log += succdssinfo
	}
}
