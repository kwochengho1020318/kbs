package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"main/services"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/tmc/langchaingo/documentloaders"
)

func InsertCsv(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	filename := string(body)
	Table_Name := r.URL.Query().Get("Table_Name")
	f, err := os.Open("upload/" + filename)
	if err != nil {
		ReturnDBError(w, err)
		return
	}
	csv := documentloaders.NewCSV(f)
	ctx := context.Background()
	documents, err := csv.Load(ctx)
	if err != nil {
		ReturnDBError(w, err)
		return
	}
	lookup, err := getlookup(documents[0].PageContent, Table_Name)
	fmt.Println("error")
	fmt.Println(lookup)
	if err != nil {
		ReturnDBError(w, err)
		return
	}
	var jsons []map[string]interface{}
	for _, document := range documents {
		rawString := document.PageContent
		rawString = strings.ReplaceAll(rawString, "\\n", "\n")

		// 使用正则表达式分割成键值对
		pattern := `(\w+): (.+?)\n`
		re := regexp.MustCompile(pattern)
		matches := re.FindAllStringSubmatch(rawString, -1)

		// 创建一个 map 存储键值对
		data := make(map[string]interface{})
		for _, match := range matches {
			key := match[1]
			value := match[2]
			data[key] = value
		}
		jsons = append(jsons, data)
	}
	jsonData, err := json.Marshal(jsons)
	if err != nil {
		ReturnDBError(w, err)
		return
	}
	result, err := datainsert(w, Table_Name, jsonData, lookup)
	if err != nil {
		ReturnDBError(w, err)
		return
	}

	services.ResponseWithText(w, 200, result)
}
