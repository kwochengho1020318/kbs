package main

import (
	"fmt"
	"main/config"
	"main/routes"
	"net/http"
	"os"
)

func main() {
	// session.InitManager(
	// 	session.SetStore(auth.NewClient()),
	// )
	myconfig := config.NewConfig()
	// fmt.Println(myconf)
	router := routes.NewRouter()
	envPort := os.Getenv("ASPNETCORE_PORT")
	hostport := fmt.Sprintf("%s:%s", myconfig.App.Url, envPort)
	http.ListenAndServeTLS(hostport, "cert.pem", "key.pem", router)

}
