package model
import (
  "time"
  jwt "github.com/dgrijalva/jwt-go"
)

//Status is the JSON struct to be returned
type Status struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Status    string `json:"status"`
}

type PaginateParams struct {
	PerPage  int    `json:"per_page"`
	CurrentPage  int   `json:"current_page"`
}
// NewPaginateParams create new instance of PaginateParams
func NewPaginateParams() PaginateParams {
   paginateParams := PaginateParams{}
   paginateParams.PerPage = 10
   paginateParams.CurrentPage = 1
   return paginateParams
}

type ResponseObj struct {
    PerPage     int
    Total   int
    CurrentPage   int
    Data interface{}
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
     AcceptingHosts []AcceptingHosts
}
type AccessRules struct {
		 AccessruleId int64 `gorm:"primary_key;json:"access_rule_id"`
		 ApplicationId int 	`json:"application_id" `
		 Email int `json:"user_id" `
		 Password int `json:"device_id"`
		 Description int `json:"group_id"`
		 Permission int `json:"access_rule_type"`
		 AcceptingHostId int `json:"description"`
		 Enabled int `json:"enabled"`
		 CreatedById int `json:"created_by_id"`
		 CreatedTime  time.Time `json:"created_time"`
		 UpdatedTime time.Time `json:"updated_time"`
}
type AcceptingHosts struct {
		 AcceptingHostId int64 `gorm:"primary_key;json:"accepting_host_id"`
		 Name string 	`json:"name"`
		 Password string `json:"password" `
		 Description string `json:"description"`
		 LastLoginTime time.Time `json:"last_login_time"`
		 Enabled int `json:"enabled"`
		 CreatedById int `json:"created_by_id"`
		 CreatedTime  time.Time `json:"created_time"`
		 UpdatedTime time.Time `json:"updated_time"`
}


type Reminder struct {
	Id        int64     `gorm:"primary_key;AUTO_INCREMENT"`
	Message   string    `sql:"size:1024" json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
}
