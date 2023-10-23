package api

import (
	"bytes"
	"encoding/json"
	"io"
	"main/config"
	"net/http"
)

func InsertHtml(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	config := config.NewConfig()
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params)

	url := params["url"]
	uid := params["uid"]
	table_name := params["table_name"]
	body := params["prompt"]

	chaturl := config.App.ChatUrl
	targeturl := chaturl + "api/chat/ExtractHtml?prompt=" + body + "&uid=" + uid
	payload := []byte(url)
	resp, err := http.Post(targeturl, "application/text", bytes.NewReader(payload))
	result, _ := io.ReadAll(resp.Body)
	if err != nil {
		ReturnDBError(w, err)
	}
	datainsert(w, table_name, result, nil)
}
