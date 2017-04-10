package controller

import (
	"encoding/json"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	_ "../db"
	_ "../model"
)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Country struct {
	Code string
	Name string
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

//GetTasksFuncAPI fetches tasks depending on the request, the authorization will be taken care by our middleare
//in this function we will return all the tasks to the user or tasks per category
//GET /api/get-tasks/

func GetTasksFuncAPI(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(
		[]Country{
			Country{
				Code: "FR",
				Name: "France",
			},
			Country{
				Code: "US",
				Name: "United States",
			},
		},
	)
}

//AddTaskFuncAPI will add the tasks for the user
func AddTaskFuncAPI(w rest.ResponseWriter, r *http.Request) {
	status:=true
	json.NewEncoder(w).Encode(status)
}

//UpdateTaskFuncAPI will add the tasks for the user
func UpdateTaskFuncAPI(w http.ResponseWriter, r *http.Request) {
	status:=true
	json.NewEncoder(w).Encode(status)
}

//DeleteTaskFuncAPI will add the tasks for the user
func DeleteTaskFuncAPI(w http.ResponseWriter, r *http.Request) {
	status:=true
	json.NewEncoder(w).Encode(status)
}

//GetDeletedTaskFuncAPI will get the deleted tasks for the user
func GetDeletedTaskFuncAPI(w http.ResponseWriter, r *http.Request) {
	status:=true
	json.NewEncoder(w).Encode(status)
}

//GetCategoryFuncAPI will return the categories for the user
//depends on the ID that we get, if we get all, then return all categories of the user
func GetCategoryFuncAPI(w http.ResponseWriter, r *http.Request) {
	status:=true
	json.NewEncoder(w).Encode(status)
}

//AddCategoryFuncAPI will add the category for the user
func AddCategoryFuncAPI(w http.ResponseWriter, r *http.Request) {
	status:=true
	json.NewEncoder(w).Encode(status)
}

//UpdateCategoryFuncAPI will update the category for the user
func UpdateCategoryFuncAPI(w http.ResponseWriter, r *http.Request) {
	status:=true
	json.NewEncoder(w).Encode(status)
}

//DeleteCategoryFuncAPI will delete the category for the user
func DeleteCategoryFuncAPI(w http.ResponseWriter, r *http.Request) {
	status:=true
	json.NewEncoder(w).Encode(status)
}
