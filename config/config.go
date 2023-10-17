package config

import (
	"encoding/json"
	"io"
	"os"
)

type BAConfig struct {
	App struct {
		Url        string
		Port       int
		LoginSite  string
		CookieName string
		ChatUrl    string
	}
	Database struct {
		Driver    string
		Server    string
		User      string
		Password  string
		Port      int
		DB_Name   string
		DockeHost string
	}
	Redis struct {
		DockerAddr string
		Addr       string
		Password   string
		Db         int
		Mode       string
	}
}

func NewConfig() BAConfig {
	confFile, err := os.Open("appsettings.json")
	if err != nil {
		panic("Unable to open config file")
	}
	defer confFile.Close()
	conf, err := io.ReadAll(confFile)
	if err != nil {
		panic("unable to read config file")
	}
	myconf := BAConfig{}
	err = json.Unmarshal(conf, &myconf)
	if err != nil {
		panic("parsing config file Error")
	}
	//fmt.Printf("%+v", myconf)
	return myconf
}
