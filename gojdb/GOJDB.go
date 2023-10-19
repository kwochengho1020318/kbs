package gojdb

import (
	"database/sql"
	"fmt"
)

type GOJDB struct {
	connection *sql.DB
	Params     *[]interface{}
}

func NewGOJDB() *GOJDB {
	jdb := &GOJDB{}
	jdb.connection, _ = SqlConnection()

	a := make([]interface{}, 0)
	jdb.Params = &a
	jdb.ParaClear()
	return jdb
}
func (db GOJDB) ParaClear() {
	*db.Params = nil
	//query:=n amedParameterQuery.NewNamedParameterQuery()
}

func (db GOJDB) ParaAdd(key string, param interface{}) {
	*db.Params = append(*db.Params, sql.Named(key, param))
}

func (db GOJDB) NonQueryWithParameters(sqlstring string) (int64, error) {
	stmt, err := db.connection.Prepare(sqlstring)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(*db.Params...)
	if err != nil {
		fmt.Println(sqlstring)
		return 0, err
	}

	rows, _ := result.RowsAffected()
	return rows, err
}

func (db GOJDB) QueryJsonWithParameters(sqlstring string) ([]interface{}, error) {
	rows, err := db.connection.Query(sqlstring, *db.Params...)

	if rows == nil {
		return make([]interface{}, 0), err
	}
	out := RowsToJson(rows)
	defer rows.Close()
	return out, err
}

func (db GOJDB) ScalarWithParameters(sqlstring string) (string, error) {
	row := db.connection.QueryRow(sqlstring, *db.Params...)
	var firstcolumnValue string
	err := row.Scan(&firstcolumnValue)

	if row == nil {
		return "", err
	}
	return firstcolumnValue, err
}

func RowsToJson(rows *sql.Rows) []interface{} {
	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	count := len(columns)
	tableData := make([]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}

		tableData = append(tableData, entry)
	}
	//jsonData, err := json.Marshal(tableData)
	// if err != nil {

	// }
	return tableData
}
