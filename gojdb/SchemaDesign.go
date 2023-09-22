package gojdb

import (
	"fmt"
	"strings"
)

func (db GOJDB) UpdateTable(table map[string]interface{}) {
	db.ParaClear()
	sqlstring := fmt.Sprintf("Select object_id from sys.objects where name = '%s'", table["TableName"])
	fmt.Println(sqlstring)
	result, err := db.Scalar(sqlstring, nil)
	if err != nil {
		panic(err)
	}
	if len(result) <= 0 {
		tableName := table["TableName"].(string)
		columns := table["Columns"].([]interface{})
		var sqlColumns []string

		for _, col := range columns {
			colData := col.(map[string]interface{})
			name := colData["name"].(string)
			typ := colData["type"].(string)
			length := colData["length"].(float64)
			sqlColumns = append(sqlColumns, fmt.Sprintf("%s %s(%d)", name, typ, int(length)))
		}

		createTableSQL := fmt.Sprintf("CREATE TABLE %s (%s);", tableName, strings.Join(sqlColumns, ", "))

		fmt.Println(createTableSQL)
	}
}
