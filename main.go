package main

import (
	"fmt"
	"main/config"
	"main/routes"
	"net/http"
)

func main() {
	// session.InitManager(
	// 	session.SetStore(auth.NewClient()),
	// )
	myconfig := config.NewConfig("appsettings.json")
	// fmt.Println(myconf)
	router := routes.NewRouter()
	hostport := fmt.Sprintf("%s:%d", myconfig.App.Url, myconfig.App.Port)

	http.ListenAndServeTLS(hostport, "cert.pem", "key.pem", router)

}
