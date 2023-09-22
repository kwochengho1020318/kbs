package gojdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func DbProtocolValid(table map[string]interface{}) (bool, error) {
	if table["TableName"] == nil {
		return false, errors.New("TableName missing")
	}

	_, ok := table["Columns"].([]interface{})
	if !ok {
		return false, errors.New("attribute columns expect array")
	}
	return true, nil
}

type Column struct {
	Name        string  `json:"name"`
	Typ         string  `json:"type"`
	Length      float64 `json:"length"`
	Notnull     int     `json:"Notnull"`
	Is_identity int     `json:"Is_Identity"`
}

func (col *Column) AddColumnString() string {
	var notnullstring string
	var identitystring string
	lengthstring := fmt.Sprintf("(%d)", int(col.Length))
	if col.Notnull == 1 {
		notnullstring = "Not Null"
	}
	if col.Is_identity == 1 {
		identitystring = "IDENTITY (1, 1)"
	}
	if col.Typ == "int" {
		lengthstring = ""
	}

	return fmt.Sprintf("%s %s%s %s %s", col.Name, col.Typ, lengthstring, identitystring, notnullstring)
}
func NewColumn(colData map[string]interface{}) (*Column, error) {
	str, err := json.Marshal(colData)
	if err != nil {
		return nil, err
	}
	var column Column
	err = json.Unmarshal(str, &column)
	if err != nil {
		return nil, err
	}
	return &column, err

}
func CompareAndUpdateColumn(column *Column, syscolumn []interface{}) {

}

func (db GOJDB) UpdateTable(table map[string]interface{}) error {
	db.ParaClear()
	sqlstring := fmt.Sprintf("Select object_id from sys.tables where name = '%s'", table["TableName"])
	fmt.Println(sqlstring)
	result, err := db.Scalar(sqlstring, nil)
	if err != nil {
		panic(err)
	}
	tableName := table["TableName"].(string)
	columns := table["Columns"].([]interface{})
	var sqlColumns []string
	if len(result) <= 0 {
		for _, col := range columns {
			colData := col.(map[string]interface{})
			column, err := NewColumn(colData)
			if err != nil {
				panic(err)
			}
			sqlColumns = append(sqlColumns, column.AddColumnString())
		}

		createTableSQL := fmt.Sprintf("CREATE TABLE %s (%s);", tableName, strings.Join(sqlColumns, ", "))
		db.NonQuery(createTableSQL, nil)

	} else {
		var newColumns []string
		var emptycolumn bool
		for _, col := range columns {
			db.ParaClear()
			colData := col.(map[string]interface{})
			column, err := NewColumn(colData)
			if err != nil {
				panic(err)
			}
			columstring := fmt.Sprintf("Select name,system_type_id,max_length,is_nullable,is_identity from sys.columns where name = '%s' and object_id = %s", column.Name, result)
			result, _ := db.QueryData(columstring, nil)
			fmt.Println(result...)
			fmt.Println(column.AddColumnString())
			if len(result) <= 0 {
				emptycolumn = true
				newColumns = append(newColumns, column.AddColumnString())
			} else {
				CompareAndUpdateColumn(column, result)
			}
		}
		if emptycolumn {
			alterString := fmt.Sprintf("Alter table %s ADD %s;", tableName, strings.Join(newColumns, ", "))
			fmt.Println(alterString)
			rowsaffected := db.NonQuery(alterString, nil)
			fmt.Println(rowsaffected)
		}

	}
	return nil
}
