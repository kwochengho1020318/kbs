package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"main/config"
	"net/http"
	"regexp"
	"strings"

	xj "github.com/basgys/goxml2json"
)

func InsertXml(w http.ResponseWriter, r *http.Request) {
	config := config.NewConfig()
	Table_Name := r.URL.Query().Get("Table_Name")
	url := config.App.ChatUrl + "api/chat/check?Table_Name=" + Table_Name
	content, err := io.ReadAll(r.Body)
	if err != nil {
		ReturnDBError(w, err)
		return
	}
	body := string(content)
	jsonstr, err := xml_to_json_string(body)
	if err != nil {
		ReturnDBError(w, err)
		return
	}
	payload := []byte(jsonstr)
	resp, err := http.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		ReturnDBError(w, err)
	}
	defer resp.Body.Close()
	var lookup map[string]interface{}
	result, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(result, &lookup)
	if err != nil {
		ReturnDBError(w, err)
		return
	}
	for key, value := range lookup {
		jsonstr = strings.Replace(jsonstr, key, value.(string), -1)
	}
	datainsert(w, Table_Name, []byte(jsonstr))

}

func xml_to_json_string(body string) (string, error) {
	if !isXMLFormat(body) {
		fmt.Println(body)
		resp, err := http.Get(body)
		if err != nil {
			println("err")
			return "", err
		}
		fmt.Println(resp)
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return "", errors.New("status not ok")
		}
		xmlData, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		fmt.Println(xmlData)
		xml := strings.NewReader(string(xmlData))
		json, err := xj.Convert(xml)
		if err != nil {
			return "", err
		}
		return json.String(), nil
	} else {
		xml := strings.NewReader(body)
		json, err := xj.Convert(xml)
		if err != nil {
			return "", err
		}
		return json.String(), nil
	}
}
func isXMLFormat(input string) bool {
	xmlRegex := regexp.MustCompile(`<\s*\w+\s*>`)
	return xmlRegex.MatchString(input)
}
