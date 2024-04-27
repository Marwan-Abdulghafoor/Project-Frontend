package main

import (
	"fmt"
	"log"
	"net/http"

	"murwan.net/fiephrs-frontend/routes"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	routes.HomepageRoutes()
	routes.RegistrationPage()
	routes.Register()
	routes.Login()
	routes.LoginPage()
	routes.Logout()
	routes.EditProfile()

	log.Print("Server is listening on port 9091 ...")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println(err.Error())
	}

}
