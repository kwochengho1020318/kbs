package gojdb

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ProcedureDefinition struct {
	ProcedureName string      `json:"procedureName"`
	Parameters    []Parameter `json:"parameters"`
	Body          string      `json:"body"`
}

type Parameter struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewStore(instore map[string]interface{}) *ProcedureDefinition {
	str, _ := json.Marshal(instore)
	var store ProcedureDefinition
	json.Unmarshal(str, &store)
	return &store
}

func (db *GOJDB) UpdateStoreProcedure(instore map[string]interface{}) {
	store := NewStore(instore)
	exsists, _ := db.Scalar(fmt.Sprintf("select name from sys.procedures where name = '%s'", store.ProcedureName), nil)
	var create_or_alter string
	if exsists == "" {
		create_or_alter = "CREATE"
	} else {
		create_or_alter = "ALTER"
	}

	var sb strings.Builder
	sb.WriteString(create_or_alter + " PROCEDURE " + store.ProcedureName + "\n")

	// Add parameters
	for i, param := range store.Parameters {
		sb.WriteString("    " + param.Name + " " + param.Type)
		if i < len(store.Parameters)-1 {
			sb.WriteString(",\n")
		}
	}
	sb.WriteString("\nAS\nBEGIN\n")
	sb.WriteString("    " + store.Body + "\nEND;\n")
	db.NonQuery(sb.String(), nil)
}
