package gojdb

import (
	"fmt"
	"strings"
)

// /////傳入sqlstring與條件 返回datatable
func (db GOJDB) QueryData(sqlstring string, params map[string][]string) ([]interface{}, error) {
	db.ParaClear()
	for key, element := range params {
		db.ParaAdd(key, element[0])
	}
	return db.QueryJsonWithParameters(sqlstring)
}

// 給定sqlstring與條件，返回string
func (db GOJDB) Scalar(sqlstring string, params map[string][]string) (string, error) {
	db.ParaClear()

	for key, element := range params {
		db.ParaAdd(key, element[0])
	}
	return db.ScalarWithParameters(sqlstring)
}

// 傳入sqlstring 與params執行NonQuery
func (db GOJDB) NonQuery(sqlstring string, params map[string][]string) int64 {
	db.ParaClear()
	for key, element := range params {
		db.ParaAdd(key, element[0])
	}
	stmt, _ := db.connection.Prepare(sqlstring)
	result, _ := stmt.Exec(*db.Params...)
	rows, _ := result.RowsAffected()
	return rows
}

// 給table,條件執行刪除
func (db GOJDB) Delete(table string, params map[string][]string) (int64, error) {
	db.ParaClear()

	sqlstring := fmt.Sprintf("Delete %s where 1=1", table)
	for key, element := range params {
		sqlstring += fmt.Sprintf("and %s = @%s", key, key)
		db.ParaAdd(key, element[0])
	}
	return db.NonQueryWithParameters(sqlstring)
}

// 給table,參數執行insert
func (db GOJDB) Insert(table string, params map[string]interface{}) (int64, error) {
	db.ParaClear()
	columns := []string{}
	values := []string{}

	for key, element := range params {
		columns = append(columns, key)
		values = append(values, "@"+key)
		db.ParaAdd(key, element)
	}
	columnsStr := strings.Join(columns, ", ")
	valuesStr := strings.Join(values, ", ")
	sqlstring := fmt.Sprintf("INSERT INTO %s(%s)values(%s)", table, columnsStr, valuesStr)
	return db.NonQueryWithParameters(sqlstring)
}

// 給table,colmns,condition執行Update
func (db GOJDB) Update(table string, params map[string]interface{}, condition map[string][]string) (int64, error) {
	db.ParaClear()
	sqlstring := fmt.Sprintf("Update %s Set ", table)
	altered := ""
	for key, element := range params {
		db.ParaAdd(key, element)
		altered += fmt.Sprintf(",%s = @%s", key, key)
	}
	altered = strings.TrimLeft(altered, " ")
	altered = strings.Replace(altered, ",", "", 1)
	sqlstring += altered
	sqlstring += " where 1=1 "
	
	for key, element := range condition {
		sqlstring += fmt.Sprintf("and %s = @%s", key, key)
		db.ParaAdd(key, element[0])
	}
	return db.NonQueryWithParameters(sqlstring)
}
