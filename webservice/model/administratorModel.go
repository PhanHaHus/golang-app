package model
import (
  "time"
  jwt "github.com/dgrijalva/jwt-go"
)
// import(
// 	"../db"
// )
/*
Package types is used to store the context struct which
is passed while templates are executed.
*/

//Status is the JSON struct to be returned
type Status struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Status    string `json:"status"`
}

type LoginParams struct {
	UserName  string    `json:"user_name"`
	Password  string   `json:"password"`
}
type SearchAdminParams struct {
	Name  string    `json:"name"`
	Description  string   `json:"description"`
}

// jwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	UserName  string `json:"user_name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type Administrators struct {
		 AdministratorId int64 `gorm:"primary_key;json:"administrator_id"`
		 Name string 	`json:"name" `
		 Email string `json:"email" `
		 Password string `json:"password"`
		 Description string `json:"description"`
		 Permission string `json:"permission"`
		 AcceptingHostId int `json:"accepting_host_id"`
		 Enabled int `json:"enabled"`
		 CreatedById int `json:"created_by_id"`
		 CreatedTime  time.Time `json:"CreatedTime"`
		 UpdatedTime time.Time `json:"UpdatedTime"`
}


type Reminder struct {
	Id        int64     `gorm:"primary_key;AUTO_INCREMENT"`
	Message   string    `sql:"size:1024" json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
}
