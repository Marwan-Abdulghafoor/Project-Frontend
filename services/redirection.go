package services

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func Redirection(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("views/index.htm"))
	tpl.Execute(w, nil)
}

func Registration(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("views/registration.htm"))
	tpl.Execute(w, nil)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("views/loginPage.htm"))
	tpl.Execute(w, nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	register(r)
	tpl := template.Must(template.ParseFiles("views/welcome.htm"))
	tpl.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	login(w, r)
}

func Profile(w http.ResponseWriter, r *http.Request) {

	tpl := template.Must(template.ParseFiles("views/profile.htm"))

	content, err := getProfileInfo(r)
	if err != nil {
		return
	}
	tpl.Execute(w, content)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("views/index.htm"))
	cookie := &http.Cookie{
		Name:   "sessionid",
		MaxAge: -1,
	}
	cookie2 := &http.Cookie{
		Name:   "user_id",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	http.SetCookie(w, cookie2)
	tpl.Execute(w, nil)
}

func EditProfile(w http.ResponseWriter, r *http.Request) {
	var (
		request  EditRequest
		response ResponseDTO
	)

	FirstName := r.FormValue("firstName")
	LastName := r.FormValue("lastName")
	Email := r.FormValue("email")
	Password := r.FormValue("password")
	Contacts := r.FormValue("contacts")
	DOB := r.FormValue("age")
	Btype := r.FormValue("btype")
	StreetAddress := r.FormValue("streetAddress")
	City := r.FormValue("city")
	State := r.FormValue("state")
	ZipCode := r.FormValue("zip")
	Gender := r.FormValue("gender")
	Allergies := r.FormValue("allergies")
	Surgeries := r.FormValue("surgeries")
	CMC := r.FormValue("cmc")
	Medications := r.FormValue("medications")
	Immunization := r.FormValue("immunization")
	FamilyHistory := r.FormValue("familyHistory")
	PhysicianName := r.FormValue("physicianName")
	PhysicianContact := r.FormValue("physicianContact")
	Emergency := r.FormValue("emergencyContact")
	IceInstructions := r.FormValue("ice_instructions")

	request.ProfileInfo.FirstName = FirstName
	request.ProfileInfo.LastName = LastName
	request.ProfileInfo.Email = Email
	request.ProfileInfo.Password = Password
	request.ProfileInfo.Dob = DOB
	request.ProfileInfo.BloodType = Btype
	request.ProfileInfo.AddressDTO.StreetAddress = StreetAddress
	request.ProfileInfo.AddressDTO.City = City
	request.ProfileInfo.AddressDTO.State = State
	request.ProfileInfo.AddressDTO.ZipCode = ZipCode
	request.ProfileInfo.Sex = Gender
	request.ProfileInfo.Allergies = Allergies
	request.ProfileInfo.Surgeries = Surgeries
	request.ProfileInfo.PhoneNumber = Contacts
	request.ProfileInfo.Chronic = CMC
	request.ProfileInfo.EmergencyContact = Emergency
	request.ProfileInfo.IceInstructions = IceInstructions
	request.ProfileInfo.Medications = Medications
	request.ProfileInfo.Immunization = Immunization
	request.ProfileInfo.FamilyHistory = FamilyHistory
	request.ProfileInfo.PhysicianName = PhysicianName
	request.ProfileInfo.PhysicianContact = PhysicianContact

	cookie, err := r.Cookie("user_id")
	if err == http.ErrNoCookie {
		return
	}
	request.Id, _ = strconv.Atoi(cookie.Value)
	req, _ := json.Marshal(request)
	URL := "http://localhost:9093/editProfile"

	err = callURL(req, URL, &response)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Profile(w, r)
}
