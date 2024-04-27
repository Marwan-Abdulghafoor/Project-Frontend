package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func register(r *http.Request) {
	var (
		request  ProfileInfo
		response AuthenticationResponseDTO
	)

	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	email := r.FormValue("email")
	password := r.FormValue("password")
	phoneNumber := r.FormValue("phoneNumber")
	emergencyContact := r.FormValue("emergency")
	emergencyIns := r.FormValue("emergencyIns")
	streetAddress := r.FormValue("streetAddress")
	city := r.FormValue("city")
	state := r.FormValue("state")
	zipCode := r.FormValue("zipCode")
	dob := r.FormValue("dob")
	sex := r.FormValue("sex")
	bloodType := r.FormValue("bloodType")
	chronic := r.FormValue("chronic")
	allergies := r.FormValue("allergies")
	surgeries := r.FormValue("surgeries")
	medications := r.FormValue("medications")
	immunization := r.FormValue("immunization")
	familyHistory := r.FormValue("familyHistory")
	physicianName := r.FormValue("physicianName")
	physicianContact := r.FormValue("physicianContact")

	request = ProfileInfo{
		FirstName:        firstName,
		LastName:         lastName,
		Email:            email,
		Password:         password,
		PhoneNumber:      phoneNumber,
		EmergencyContact: emergencyContact,
		IceInstructions:  emergencyIns,
		AddressDTO: AddressDTO{
			StreetAddress: streetAddress,
			City:          city,
			State:         state,
			ZipCode:       zipCode,
		},
		Dob:              dob,
		Sex:              sex,
		BloodType:        bloodType,
		Chronic:          chronic,
		Allergies:        allergies,
		Surgeries:        surgeries,
		Medications:      medications,
		Immunization:     immunization,
		FamilyHistory:    familyHistory,
		PhysicianName:    physicianName,
		PhysicianContact: physicianContact,
	}

	req, _ := json.Marshal(request)

	URL := "http://localhost:8080/api/v1/user/auth/register"

	err := callURL(req, URL, &response)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("email")
	password := r.FormValue("password")
	println("&&&&&", username, password)
	var (
		request  LoginRequest
		response AuthenticationResponseDTO
	)
	request.Email = username
	request.Password = password

	req, _ := json.Marshal(request)
	URL := "http://localhost:8080/api/v1/user/auth/login"
	err := callURL(req, URL, &response)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	info, _ := decryptJWT(response.Token)
	println("&&&&&2222")
	if info.Id != "" {
		cookie := &http.Cookie{
			Name:  "sessionid",
			Value: info.Token,
		}
		cookie2 := &http.Cookie{
			Name:  "user_id",
			Value: info.Id,
		}
		http.SetCookie(w, cookie)
		http.SetCookie(w, cookie2)
		println("&&&&&33333")
		Profile(w, r)
		return
	} else {
		LoginPage(w, r)
	}
}

func getProfileInfo(r *http.Request) (ProfileInfo, error) {
	var profile ProfileInfo
	var (
		request  StandardIdRequest
		response ResponseDTO
	)

	cookie, err := r.Cookie("user_id")
	if err == http.ErrNoCookie {
		return profile, err
	}

	URL := "http://localhost:9093/getProfileInfo"
	request.Id, _ = strconv.Atoi(cookie.Value)
	req, _ := json.Marshal(request)

	err = callURL(req, URL, &response)

	if err != nil {
		fmt.Println(err.Error())
		return profile, err
	}

	profile = response.ProfileInfo
	return profile, nil
}
