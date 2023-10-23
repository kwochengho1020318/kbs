package api

import (
	"bytes"
	"fmt"
	"io"
	"main/config"
	"net/http"
)

func InsertHtml(w http.ResponseWriter, r *http.Request) {
	config := config.NewConfig()
	prompt := r.URL.Query().Get("prompt")
	uid := r.URL.Query().Get("uid")
	table_name := r.URL.Query().Get("table_name")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	chaturl := config.App.ChatUrl
	targeturl := chaturl + "api/chat/ExtractHtml?prompt" + prompt + "&uid=" + uid
	payload := body
	resp, err := http.Post(targeturl, "application/text", bytes.NewReader(payload))
	result, _ := io.ReadAll(resp.Body)
	if err != nil {
		ReturnDBError(w, err)
	}
	datainsert(w, table_name, result, nil)
}
