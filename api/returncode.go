package api

import (
	"main/services"
	"net/http"
	"strings"
)

type ErrorCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ReturnDBError(w http.ResponseWriter, err error) {

	switch {
	//DB not connected
	case strings.Contains(err.Error(), "connection"):
		services.ResponseWithJson(w, http.StatusInternalServerError, ErrorCode{001, "DB not connected"})
	//syntax error
	case strings.Contains(err.Error(), "syntax"):
		services.ResponseWithJson(w, http.StatusInternalServerError, ErrorCode{002, "syntax error"})
	//invalid column name
	case strings.Contains(err.Error(), "Invalid column name"):
		services.ResponseWithJson(w, http.StatusBadRequest, ErrorCode{003, "Invalid column name"})
	case strings.Contains(err.Error(), "duplicate"):
		services.ResponseWithJson(w, http.StatusBadRequest, ErrorCode{004, "重複的 key"})
	default:
		services.ResponseWithJson(w, http.StatusInternalServerError, ErrorCode{999, "unknown"})
	}
}
