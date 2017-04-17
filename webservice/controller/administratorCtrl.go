package controller
import (
  "github.com/ant0ine/go-json-rest/rest"
  "time"
  _ "log"
  "net/http"
	database "../db"
)

type Reminder struct {
	Id        int64     `gorm:"primary_key;AUTO_INCREMENT"`
	Message   string    `sql:"size:1024" json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
}

func GetAllReminders(w rest.ResponseWriter, r *rest.Request) {
  tx := database.MysqlConn().Begin()
  reminders := []Reminder{}
	tx.Find(&reminders)
	w.WriteJson(&reminders)
}

func GetReminder(w rest.ResponseWriter, r *rest.Request) {
  tx := database.MysqlConn().Begin()
	id := r.PathParam("id")
	reminder := Reminder{}
	if tx.First(&reminder, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&reminder)
}


func PostReminder(w rest.ResponseWriter, r *rest.Request) {
  tx:= database.MysqlConn().Begin()
	reminder := Reminder{}
	if err := r.DecodeJsonPayload(&reminder); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  data:= Reminder{
    Message: reminder.Message,
  }

	if err := tx.Create(&data).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  tx.Commit()
	w.WriteJson(&data)
}


func  PutReminder(w rest.ResponseWriter, r *rest.Request) {
  tx:= database.MysqlConn().Begin()
	id := r.PathParam("id")
	reminder := Reminder{}
	if tx.First(&reminder, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Reminder{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reminder.Message = updated.Message

	if err := tx.Save(&reminder).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
  tx.Commit()
	w.WriteJson(&reminder)
}

func  DeleteReminder(w rest.ResponseWriter, r *rest.Request) {
  tx:= database.MysqlConn().Begin()
	id := r.PathParam("id")
	reminder := Reminder{}
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
