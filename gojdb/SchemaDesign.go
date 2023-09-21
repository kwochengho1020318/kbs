package gojdb

import (
	"fmt"
)

func (db GOJDB) UpdateTable(table map[string]interface{}) {

	sqlstring := fmt.Sprintf("Select object_id from sys.objects where name = %s", table["TableName"])
	result, err := db.Scalar(sqlstring, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
