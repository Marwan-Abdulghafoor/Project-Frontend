package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func callURL(req []byte, URL string, res interface{}) (err error) {
	responseBody := bytes.NewBuffer(req)
	response, err := http.Post(URL, "application/json", responseBody)
	if err != nil {
		fmt.Println("WERE")
		fmt.Println(err.Error())
		return
	}
	data, _ := ioutil.ReadAll(response.Body)
	println(data)
	err = json.Unmarshal(data, res)
	if err != nil {
		fmt.Println("qwer")
		fmt.Println(err.Error())
		return
	}
	return
}

func decryptJWT(tokenString string) (UserCoockieInfo, error) {
	var info UserCoockieInfo
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		fmt.Println(err.Error())
		return info, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		info.Id = claims["jti"].(string)
		info.Token = tokenString
	}

	if info.Id == "" {
		return info, fmt.Errorf("invalid token payload")
	}

	return info, nil
}
