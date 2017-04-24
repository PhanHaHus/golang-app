package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
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
func MiddlewareJWT(c echo.Context) (err error){
	// token := r.Header.Get("Authorization")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNDkyNjcyNDQxfQ.xevIvEkT1S2frmYutds-_Sote3EtfX6ZmqOcRrEybpk"
	IsTokenValid, username := ValidateToken(token)
	//When the token is not valid show the default error JSON document
	if !IsTokenValid {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": "Something went wrong with signing token, Authentication failed!","status":"false"})
	}
	log.Println("token is valid " + username + " is logged in")
	return c.JSON(http.StatusOK,map[string]string{"Message": "ok logged"})
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
func LoginCtrl(c echo.Context) (err error) {
  loginParams := model.LoginParams{}
	log.Println(c.Bind(&loginParams))
	if err = c.Bind(&loginParams); err != nil {
		 return c.JSON(http.StatusInternalServerError, map[string]string{"Message": "Cant get Params","status":"false"})
	}
  username := loginParams.UserName
	password := loginParams.Password
	log.Println(username, " ", password)

	if(ValidUser(username,password)){
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Snow"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}
