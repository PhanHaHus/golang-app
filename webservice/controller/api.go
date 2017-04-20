package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"log"
	"time"
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
//ValidUser will check if the user exists in db and if exists if the username password
//combination is valid
func ValidUser(username, password string) bool {
	//If the password matches, return true
	if (username=="admin" && password == "admin") {
		return true
	}
	//by default return false
	return false
}
// Login API
func LoginCtrl(w rest.ResponseWriter, r *rest.Request) {
  loginParams := model.LoginParams{}
  if err := r.DecodeJsonPayload(&loginParams); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  username := loginParams.UserName
	password := loginParams.Password
	log.Println(username, " ", password)
	if(ValidUser(username,password)){
		// Create the Claims
		claims := MyCustomClaims{
			username,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		/* Sign the token with our secret */
		tokenString, err := token.SignedString(mySigningKey)
		if err != nil {
			log.Println("Something went wrong with signing token")
			w.WriteJson(map[string]string{"Message": "Something went wrong with signing token, Authentication failed!","status":"false"})
			return
		}
		/* Finally, write the token to the browser window */
		w.WriteJson(map[string]string{"token":tokenString,"status":"true","Message": "Authentication Success!"} )
	} else {
		w.WriteJson(map[string]string{"Message": "Authentication failed!","status":"false"})
	}

}
