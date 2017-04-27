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

	if err = c.Bind(&loginParams); err != nil {
		 return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: "Cant get Params",Status:"false"})
	}
  username := loginParams.UserName
	password := loginParams.Password
	log.Println(username, " ", password)
	if(ValidUser(username,password)){
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = username
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
			"username": username,
		})
	}
	return c.JSON(http.StatusOK, model.Status{StatusCode: http.StatusOK, Message: "Login fail, wrong username or password !",Status:"false"});
}

func SearchAdminCtrl(c echo.Context) (err error) {
  params := model.SearchAdminParams{}
  c.Bind(&params)

	log.Println("params:")
	// name:=params.Name
	// description:=params.Description
  log.Println(params)
  log.Println("------------")

	// tx:= database.MysqlConn().Begin()
	// if err := tx.First(&reminder, id).Error; err != nil {
	// 	return c.JSON(http.StatusNotFound,model.Status{StatusCode: http.StatusNotFound, Message: err.Error(),Status:"false"})
	// }
	//
	// tx.Commit()
	return c.JSON(http.StatusOK, params)
}
