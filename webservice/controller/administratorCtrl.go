package controller
import (
  "github.com/ant0ine/go-json-rest/rest"
  _ "log"
  "net/http"
	database "../db"
  model "../model"
)

// type Reminder struct {
// 	Id        int64     `gorm:"primary_key;AUTO_INCREMENT"`
// 	Message   string    `sql:"size:1024" json:"message"`
// 	CreatedAt time.Time `json:"createdAt"`
// 	UpdatedAt time.Time `json:"updatedAt"`
// 	DeletedAt *time.Time `json:"-"`
// }

func GetAllAdmin(w rest.ResponseWriter, r *rest.Request) {
  tx := database.MysqlConn().Begin()
  administrators := []model.Administrators{}
	tx.Find(&administrators)
	w.WriteJson(&administrators)
}

func GetAdminById(w rest.ResponseWriter, r *rest.Request) {
  tx := database.MysqlConn().Begin()
	id := r.PathParam("id")
	reminder := model.Administrators{}
	if tx.First(&reminder, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&reminder)
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
