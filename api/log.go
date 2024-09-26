package api

import (
	"main/config"
	"os"
)

func WriteLog(input string) error {
	config := config.NewConfig()
	f, err := os.OpenFile(config.App.LogPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {

		return err
	}

	defer f.Close()
	_, err = f.WriteString(input)
	if err != nil {
		return err
	}
	return nil
}
