package api

import (
	"fmt"
	"io"
	"main/config"
	"main/services"
	"net/http"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		WriteLog(err.Error())
		services.ResponseWithJson(w, 500, ErrorCode{Code: 1, Message: "讀取請求錯誤"})
		return
	}
	defer file.Close()
	config := config.NewConfig()

	f, err := os.OpenFile(fmt.Sprintf("%s/%s", config.App.UploadFolder, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		WriteLog(err.Error())
		services.ResponseWithJson(w, 500, ErrorCode{Code: 1, Message: err.Error()})
		return
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		WriteLog(err.Error())
		services.ResponseWithJson(w, 500, ErrorCode{Code: 1, Message: "寫入錯誤"})
		return
	}
	services.ResponseWithJson(w, 200, ErrorCode{Code: 1, Message: "success"})
}
