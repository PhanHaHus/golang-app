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
//response object
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
		 AcceptingHostId int64
		 Enabled int `json:"enabled"`
		 CreatedById int `json:"created_by_id"`
		 CreatedTime  time.Time `json:"CreatedTime"`
		 UpdatedTime time.Time `json:"UpdatedTime"`
     AcceptingHost AcceptingHosts `gorm:"ForeignKey:AcceptingHostId;AssociationForeignKey:AcceptingHostId"` // belong to AcceptingHost
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

type AccessRules struct {
		 AccessRuleId int64 `gorm:"primary_key;json:"access_rule_id"`
		 ApplicationId int64
		 UserId int `json:"user_id" `
		 DeviceId int `json:"device_id"`
		 GroupId int `json:"group_id"`
		 AccessRuleType int `json:"access_rule_type"`
		 Description int `json:"description"`
		 Enabled int `json:"enabled"`
		 CreatedById int `json:"created_by_id"`
		 CreatedTime  time.Time `json:"created_time"`
		 UpdatedTime time.Time `json:"updated_time"`
     Application Applications `gorm:"ForeignKey:ApplicationId;AssociationForeignKey:ApplicationId"` // belong to application
}

type Applications struct {
		 ApplicationId int64 `gorm:"primary_key;json:"application_id"`
		 Name string 	`json:"name" `
		 ApplicationType int `json:"application_type" `
		 AcceptingHostId int `json:"accepting_host_id"`
		 Ip int `json:"ip"`
		 Port int `json:"port"`
		 HostName string `json:"host_name"`
		 IsValidUserRequired int `json:"is_valid_user_required"`
		 IsValidDeviceRequired int `json:"is_valid_device_required"`
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
