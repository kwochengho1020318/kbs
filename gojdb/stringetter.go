package gojdb

import (
	"fmt"
	"strconv"
	"strings"
)

func (db GOJDB) SelectTableString(columns []string, table string, params map[string][]string) string {
	columnstr := ""
	if columns == nil {
		columnstr = "*"
	} else {
		for _, element := range columns {
			columnstr += element + ","
		}
	}
	condition := ""
	pagesize := 100
	page := 1
	paging := ""
	ispage := false
	for key, element := range params {
		if strings.ToLower(key) == "pagesize" {
			num, _ := strconv.Atoi(element[0])
			pagesize = num
			ispage = true
			continue
		}
		if strings.ToLower(key) == "page" {
			page, _ = strconv.Atoi(element[0])
			ispage = true
			continue
		}
		condition += fmt.Sprintf(" and %s = @%s", key, key)
	}
	if ispage {
		paging += fmt.Sprintf("ORDER BY (Select Null) OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", (page-1)*pagesize, pagesize)

	}
	sqlstring := fmt.Sprintf("SELECT %sfrom %s where 1=1 %s %s", columnstr, table, condition, paging)
	return sqlstring
}
func (db GOJDB) ScalarString(colmn string, table string, params map[string][]string) string {
	sqlstring := fmt.Sprintf("Select  %s from %s where 1=1", colmn, table)
	for key := range params {
		sqlstring += fmt.Sprintf(" and %s = @%s", key, key)
	}
	return sqlstring
}
