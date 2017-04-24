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
  MiddlewareJWT(w,r)
  tx := database.MysqlConn().Begin()
  administrators := []model.Administrators{}
	tx.Order("administrators.administrator_id desc").Limit(10).Find(&administrators)
  tx.Commit()
	w.WriteJson(&administrators)
}

func SearchCtrl(w rest.ResponseWriter, r *rest.Request) {
  tet := r
  // limit := r.Form.Get("limit")
  log.Println("params:")
  log.Println(tet)
  log.Println("------------")

	w.WriteHeader(http.StatusOK)
}

func GetAdminById(w rest.ResponseWriter, r *rest.Request) {
  log.Println("GetAdminById")
  tx := database.MysqlConn().Begin()
	administratorsId := r.PathParam("id")
  log.Println("Id: ",administratorsId)
	administrator := model.Administrators{}
	if tx.First(&administrator, administratorsId).Error != nil {
		rest.NotFound(w, r)
		return
	}
  tx.Commit()
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


func PutAdmin(w rest.ResponseWriter, r *rest.Request) {
  tx:= database.MysqlConn().Begin()
	administrator_id := r.PathParam("id")

	administrator := model.Administrators{}
	if tx.First(&administrator, administrator_id).Error != nil {
		rest.NotFound(w, r)
		return
	}
  log.Println("id:")
  log.Println(tx)

	updated := model.Administrators{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	administrator.Name = updated.Name
	administrator.Email = updated.Email
	administrator.Password = updated.Password
	administrator.Description = updated.Description
	administrator.AcceptingHostId = updated.AcceptingHostId
	administrator.Enabled = updated.Enabled
	administrator.CreatedById = updated.CreatedById

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
