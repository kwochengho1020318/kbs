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
	Name          string  `json:"name"`
	Typ           string  `json:"type"`
	Length        float64 `json:"length"`
	Is_nullable   bool    `json:"Is_nullable"`
	Is_identity   bool    `json:"Is_Identity"`
	Default_Value string  `json:"Default_Value"`
	ForeignKey    string  `json:"ForeignKey"`
}
type View struct {
	ViewName       string
	SelectColumns  []string
	FromTables     []string
	WhereCondition string
	JoinTable      string
	JoinCondition  string
	JoinType       string
}

func NewView(view map[string]interface{}) *View {
	viewstr, _ := json.Marshal(view)
	var temp View
	json.Unmarshal(viewstr, &temp)
	fmt.Println(temp)
	return &temp
}
func (db *GOJDB) UpdateView(inview map[string]interface{}) error {
	view := NewView(inview)
	var ViewCreateString string
	if view.WhereCondition != "" {
		ViewCreateString = fmt.Sprintf("CREATE OR ALTER VIEW %s AS\nSELECT %s\nFROM %s\n Where %s;",
			view.ViewName,
			strings.Join(view.SelectColumns, ", "),
			strings.Join(view.FromTables, ", "),
			view.WhereCondition,
		)
	} else {
		ViewCreateString = fmt.Sprintf("CREATE OR ALTER VIEW %s AS\nSELECT %s\nFROM %s\n %s JOIN %s ON %s;",
			view.ViewName,
			strings.Join(view.SelectColumns, ", "),
			strings.Join(view.FromTables, ", "),
			view.JoinType,
			view.JoinTable,
			view.JoinCondition)
	}

	fmt.Println(ViewCreateString)
	_, err := db.NonQuery(ViewCreateString, nil)
	if err != nil {
		return err
	}
	return nil
}

func (col *Column) AddColumnString() string {
	var notnullstring string
	var identitystring string
	lengthstring := fmt.Sprintf("(%d)", int(col.Length))
	if !col.Is_nullable {

		notnullstring = "Not Null"
	}
	if col.Is_identity {
		identitystring = "IDENTITY (1, 1) Primary Key "

	}

	if col.Length == 0 {
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
	column.Is_nullable = true
	err = json.Unmarshal(str, &column)
	if column.Typ == "" {
		column.Typ = "nvarchar"
	}

	if err != nil {
		return nil, err
	}
	return &column, err

}
func (db *GOJDB) UpdateColumn(column *Column, syscolumn interface{}, tableName string) error {
	//修改型別
	var typestring string
	var notnullstring string
	if !column.Is_nullable {
		sqlstring := fmt.Sprintf("Update [%s] set %s = '%s' where %s is null", tableName, column.Name, column.Default_Value, column.Name)
		_, err := db.NonQuery(sqlstring, nil)
		if err != nil {
			fmt.Println(sqlstring)
			fmt.Println(err)
			return err
		}
		notnullstring = "NOT NULL"
	}
	if column.ForeignKey != "" {
		sqlstring := fmt.Sprintf("select name from sys.key_constraints where name = 'FK_%s_%s'", tableName, column.Name)
		result, _ := db.ScalarWithParameters(sqlstring)

		if result == "" {
			sqlstring := fmt.Sprintf("ALTER TABLE %s ADD CONSTRAINT FK_%s_%s FOREGIGN KEY (%s) REFERENCES %s ", tableName, tableName, column.Name, column.Name, column.ForeignKey)
			db.NonQueryWithParameters(sqlstring)
		}
	}
	//int vs nvarchar(50)
	if column.Length != 0 {
		typestring = fmt.Sprintf("%s(%d)", column.Typ, int(column.Length))
	} else {
		typestring = column.Typ
	}
	overwritestring := fmt.Sprintf("alter table [%s] \nalter column \n %s %s %s", tableName, column.Name, typestring, notnullstring)
	_, err := db.NonQuery(overwritestring, nil)
	if err != nil {
		fmt.Println(overwritestring)
		fmt.Println(err)
		return err
	}
	//修改default值
	if column.Default_Value != "" {
		sqlstring := fmt.Sprintf("Alter Table [%s] ADD CONSTRAINT DF_%s_%s default %s FOR %s", tableName, tableName, column.Name, column.Default_Value, column.Name)
		fmt.Println(sqlstring)
		_, err := db.NonQuery(sqlstring, nil)
		if err != nil {

			return err
		}
	}

	return nil
}

func (db GOJDB) UpdateTable(table map[string]interface{}) error {
	db.ParaClear()
	updatecolumerr := ""
	sqlstring := fmt.Sprintf("Select object_id from sys.tables where name = '%s'", table["TableName"])
	result, _ := db.Scalar(sqlstring, nil)
	tableName := table["TableName"].(string)
	columns := table["Columns"].([]interface{})
	var sqlColumns []string
	//若不存在->新增table
	//fmt.Println(result)
	if result == "" {
		var constraintkey []string
		for _, col := range columns {
			colData := col.(map[string]interface{})
			column, err := NewColumn(colData)
			if err != nil {
				panic(err)
			}

			sqlColumns = append(sqlColumns, column.AddColumnString())
			if column.ForeignKey != "" {

				constraintkey = append(constraintkey, fmt.Sprintf("CONSTRAINT FK_%s_%s FOREIGN KEY (%s) REFERENCES %s on update cascade on delete cascade ", tableName, column.Name, column.Name, column.ForeignKey))
			}
		}

		createTableSQL := fmt.Sprintf("CREATE TABLE [%s] (%s %s);", tableName, strings.Join(sqlColumns, ", "), strings.Join(constraintkey, ", "))
		fmt.Println(createTableSQL)
		_, err := db.NonQuery(createTableSQL, nil)
		if err != nil {
			fmt.Println(createTableSQL)
			return err
		}

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
			//若不存在，新增欄位
			if len(result) <= 0 {
				emptycolumn = true
				newColumns = append(newColumns, column.AddColumnString())
			} else {
				err := db.UpdateColumn(column, result[0], tableName)
				if err != nil {
					temp := fmt.Sprintf("Updating table %s column %s failed", tableName, column.Name)
					fmt.Println(err)
					fmt.Println(temp)
					updatecolumerr += temp + " \n"
				}
			}
		}
		if emptycolumn {
			alterString := fmt.Sprintf("Alter table [%s] ADD %s;", tableName, strings.Join(newColumns, ", "))
			rowsaffected, err := db.NonQuery(alterString, nil)
			if err != nil {
				fmt.Println(alterString)
				return err
			}
			fmt.Println(rowsaffected)
		}

	}
	if updatecolumerr == "" {
		return nil
	} else {
		return errors.New(updatecolumerr)
	}
}
