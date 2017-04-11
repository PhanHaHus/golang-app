package model
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
}

type Administrators struct {
		 AdministratorId int64 `json:"administratorId"`
		 Name string 	`json:"name"`
		 Email string `json:"email"`
		 Password string `json:"password"`
		 Description string `json:"description"`
		 Permission string `json:"permission"`
		 AcceptingHostId int `json:"accepting_host_id"`
		 Enabled int `json:"enabled"`
		 CreatedById int `json:"created_by_id"`
		 CreatedAt  string `json:"created_time"`
		 UpdatedAt  string `json:"updated_time"`
}

// func GetAdministrator() ([]User, error) {
// 	var (
// 		users []User
// 		err   error
// 	)
//
// 	tx := gorm.MysqlConn().Begin()
// 	if err = tx.Find(&users).Error; err != nil {
// 		tx.Rollback()
// 		return users, err
// 	}
// 	tx.Commit()
//
// 	return users, err
// }
