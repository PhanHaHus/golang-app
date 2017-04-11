package controller
import (
  "github.com/ant0ine/go-json-rest/rest"
	database "../db"
)

type Reminder struct {
	Id        int64     `json:"id"`
	Message   string    `sql:"size:1024" json:"message"`
	// CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt time.Time `json:"updatedAt"`
	// DeletedAt time.Time `json:"-"`
}

func GetAllReminders(w rest.ResponseWriter, r *rest.Request) {
	reminders := []Reminder{}
  tx := database.MysqlConn().Begin()
	tx.DB.Find(&reminders)
	w.WriteJson(&reminders)
}
