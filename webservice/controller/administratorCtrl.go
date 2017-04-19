package controller
import (
  "github.com/ant0ine/go-json-rest/rest"
    "log"
  "net/http"
  _ "encoding/json"
	database "../db"
  model "../model"
)

func GetAllAdmin(w rest.ResponseWriter, r *rest.Request) {
  tx := database.MysqlConn().Begin()
  administrators := []model.Administrators{}
	tx.Find(&administrators)
	w.WriteJson(&administrators)
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

  w.WriteJson(&loginParams)
}

func GetAdminById(w rest.ResponseWriter, r *rest.Request) {
  tx := database.MysqlConn().Begin()
	id := r.PathParam("id")
	administrator := model.Administrators{}
	if tx.First(&administrator, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&administrator)
}


func PostAdmin(w rest.ResponseWriter, r *rest.Request) {
  tx:= database.MysqlConn().Begin()
	administrators := model.Administrators{}
	if err := r.DecodeJsonPayload(&administrators); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tx.Create(&administrators).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  tx.Commit()
	w.WriteJson(&administrators)
}


func  PutAdmin(w rest.ResponseWriter, r *rest.Request) {
  tx:= database.MysqlConn().Begin()
	id := r.PathParam("id")
	administrator := model.Administrators{}
	if tx.First(&administrator, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := model.Administrators{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	administrator.Name = updated.Name

	if err := tx.Save(&administrator).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
  tx.Commit()
	w.WriteJson(&administrator)
}

func  DeleteAdmin(w rest.ResponseWriter, r *rest.Request) {
  tx:= database.MysqlConn().Begin()
	id := r.PathParam("id")
	reminder := model.Administrators{}
	if tx.First(&reminder, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := tx.Delete(&reminder).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
  tx.Commit()
	w.WriteHeader(http.StatusOK)
}
