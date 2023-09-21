package gojdb

import (
	"fmt"
)

func (db GOJDB) UpdateTable(table map[string]interface{}) {
	db.ParaClear()
	sqlstring := fmt.Sprintf("Select object_id from sys.objects where name = %s", table["TableName"])
	fmt.Println(sqlstring)
	result, err := db.Scalar(sqlstring, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
