package controller

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"log"
	"encoding/json"
	model "../model"
)

type MyCustomClaims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}
var mySigningKey = []byte("secret")

//ValidateToken will validate the token
func ValidateToken(myToken string) (bool, string) {
	token, err := jwt.ParseWithClaims(myToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})
	if err != nil {
		return false, ""
	}
	claims := token.Claims.(*MyCustomClaims)
	return token.Valid, claims.Username
}
// using with request api
func MiddlewareJWT(w http.ResponseWriter, r *http.Request){
	var err error
	var status model.Status
	var message string

	token := r.Header["Token"][0]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	IsTokenValid, username := ValidateToken(token)
	//When the token is not valid show the default error JSON document
	if !IsTokenValid {
		status = model.Status{StatusCode: http.StatusInternalServerError, Message: message}
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(status)
		if err != nil {
			panic(err)
		}
		return
	}
	log.Println("token is valid " + username + " is logged in")
}
