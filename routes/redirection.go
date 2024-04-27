package routes

import (
	"net/http"

	"murwan.net/fiephrs-frontend/services"
)

func HomepageRoutes() {
	http.HandleFunc("/", services.Redirection)
}

func RegistrationPage() {
	http.HandleFunc("/registration", services.Registration)
}

func LoginPage() {
	http.HandleFunc("/loginPage", services.LoginPage)
}

func Register() {
	http.HandleFunc("/register", services.Register)
}

func Login() {
	http.HandleFunc("/login", services.Login)
}

func Logout() {
	http.HandleFunc("/logout", services.Logout)
}

func EditProfile() {
	http.HandleFunc("/editProfile", services.EditProfile)

}
