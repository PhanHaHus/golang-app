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

	auth := c.Request().Header.Get("")
	log.Println("token is  " + auth)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
	IsTokenValid, username := ValidateToken(token)
	//When the token is not valid show the default error JSON document
	if !IsTokenValid {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": "Something went wrong with signing token, Authentication failed!","status":"false"})
	}
	return c.JSON(http.StatusOK,map[string]string{"Message": username+ "ok logged"})
}
//ValidUser will check if the user exists in db and if exists if the username password
func ValidUser(username, password string) bool {
	if (username=="admin" && password == "admin") {
		return true
	}
	return false
}

// Login API
func LoginCtrl(c echo.Context) (err error) {
  loginParams := model.LoginParams{}

	// if err = c.Bind(&loginParams); err != nil {
	// 	 return c.JSON(http.StatusInternalServerError, map[string]string{"Message": "Cant get Params","status":"false"})
	// }

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


func SearchCtrl(c echo.Context) (err error) {
  params := model.SearchParams{}
  if err = c.Bind(&params); err != nil {
     return c.JSON(http.StatusInternalServerError, map[string]string{"Message": "InternalServerError","status":"false"})
  }
  log.Println("params:")
  log.Println(params)
  log.Println("------------")

	return c.JSON(http.StatusOK, params)
}
