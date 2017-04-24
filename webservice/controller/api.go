package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"log"
	"time"
	_ "encoding/json"
	model "../model"
)

type MyCustomClaims struct {
    UserName string `json:"user_name"`
    jwt.StandardClaims
}

var mySigningKey = []byte("secret")

//ValidateToken will validate the token
func ValidateToken(myToken string) (bool, string) {
log.Println("checking");
log.Println(myToken);
	token, err := jwt.ParseWithClaims(myToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})
	if err != nil {
		return false, ""
	}
	claims := token.Claims.(*MyCustomClaims)
	log.Println(token.Valid);
	return token.Valid, claims.UserName
}
// using with request api
func MiddlewareJWT(w rest.ResponseWriter, r *rest.Request){
	token := r.Header.Get("Authorization")
	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNDkyNjcyNDQxfQ.xevIvEkT1S2frmYutds-_Sote3EtfX6ZmqOcRrEybpk"
	IsTokenValid, username := ValidateToken(token)
	//When the token is not valid show the default error JSON document
	if !IsTokenValid {
		w.WriteJson(map[string]string{"Message": "Something went wrong with signing token, Authentication failed!","status":"false"})
		return
	}
	log.Println("token is valid " + username + " is logged in")
	w.WriteJson(map[string]string{"Message": "ts"})
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
