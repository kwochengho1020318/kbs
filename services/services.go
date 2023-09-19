package services

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ResponseWithText(w http.ResponseWriter, code int, payload string) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(code)
	w.Write([]byte(payload))
}
func ResponseWithHtml(w http.ResponseWriter, code int, payload []byte) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	w.Write(payload)
}
