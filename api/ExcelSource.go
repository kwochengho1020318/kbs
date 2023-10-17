package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/config"
	"main/services"
	"net/http"
	"strings"

	"github.com/xuri/excelize/v2"
)

func InsertExcel(w http.ResponseWriter, r *http.Request) {
	filename, _ := io.ReadAll(r.Body)
	Table_name := r.URL.Query().Get("Table_Name")
	if Table_name == "" {
		services.ResponseWithText(w, 400, "Table_Name not provided ")
	}
	Sheets := XlsxToJson(string(filename))
	for _, v := range Sheets {
		err := InsertSheet(w, v, Table_name)
		if err != nil {
			ReturnDBError(w, err)
			return
		}
	}
}

func InsertSheet(w http.ResponseWriter, sheet string, Table_Name string) error {
	config := config.NewConfig()
	url := config.App.ChatUrl + "api/chat/check?Table_Name=" + Table_Name
	payload := []byte(sheet)
	resp, err := http.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var lookup map[string]interface{}
	result, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(result, &lookup)
	if err != nil {

		return err
	}
	for key, value := range lookup {
		sheet = strings.Replace(sheet, key, value.(string), -1)
	}
	datainsert(w, Table_Name, []byte(sheet))
	return nil
}

func XlsxToJson(filename string) []string {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheets := f.GetSheetList()
	var Jsons []string
	for _, sheetName := range sheets {
		d, err := f.GetRows(sheetName)
		if err != nil {
			fmt.Println("error reading sheet", sheetName, ":", err)
			return nil
		}

		JsonO := ExcelRowsToJson(d)
		jsonStr, err := json.Marshal(JsonO)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		Jsons = append(Jsons, string(jsonStr))
	}
	return Jsons
}

func ExcelRowsToJson(rows [][]string) []map[string]string {
	data := make([]map[string]string, len(rows)-1)
	headers := rows[0]
	// excluding header row
	for i, row := range rows[1:] {
		data[i] = make(map[string]string)
		for j, cellValue := range row {
			data[i][headers[j]] = cellValue
		}
	}

	return data
}
