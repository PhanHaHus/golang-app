package views

/*Holds the fetch related view handlers*/

import (
	"html/template"
		 "log"
	"net/http"
	_ "time"

	// _ "github.com/thewhitetulip/Tasks/sessions"
)

var homeTemplate *template.Template
var deletedTemplate *template.Template
var completedTemplate *template.Template
var editTemplate *template.Template
var searchTemplate *template.Template
var templates *template.Template
var loginTemplate *template.Template
var addReminderTemplate *template.Template
var detailReminderTemplate *template.Template
var editReminderTemplate *template.Template


var message string //message will store the message to be shown as notification
var err error

//ShowAllTasksFunc is used to handle the "/" URL which is the default ons
//TODO add http404 error

func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("Show");
	if r.Method == "GET" {
			homeTemplate.Execute(w,nil)
		}
}
func EditReminderFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("Edit");
	if r.Method == "GET" {
			editReminderTemplate.Execute(w,nil)
		}
}

func DetailReminderFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("detail");
	if r.Method == "GET" {
			detailReminderTemplate.Execute(w,nil)
	}
}

func AddReminder(w http.ResponseWriter, r *http.Request) {
	log.Println("add");
	if r.Method == "GET" {
			addReminderTemplate.Execute(w,nil)
	}
}


// //ShowTrashTaskFunc is used to handle the "/trash" URL which is used to show the deleted tasks
// func ShowTrashTaskFunc(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		username := sessions.GetCurrentUserName(r)
// 		categories := db.GetCategories(username)
// 		context, err := db.GetTasks(username, "deleted", "")
// 		context.Categories = categories
// 		if err != nil {
// 			http.Redirect(w, r, "/trash", http.StatusInternalServerError)
// 		}
// 		if message != "" {
// 			context.Message = message
// 			message = ""
// 		}
// 		err = deletedTemplate.Execute(w, context)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

// //ShowCompleteTasksFunc is used to populate the "/completed/" URL
// func ShowCompleteTasksFunc(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		username := sessions.GetCurrentUserName(r)
// 		categories := db.GetCategories(username)
// 		context, err := db.GetTasks(username, "completed", "")
// 		context.Categories = categories
// 		if err != nil {
// 			http.Redirect(w, r, "/completed", http.StatusInternalServerError)
// 		}
// 		completedTemplate.Execute(w, context)
// 	}
// }

// //ShowCategoryFunc will populate the /category/<id> URL which shows all the tasks related
// // to that particular category
// func ShowCategoryFunc(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" && sessions.IsLoggedIn(r) {
// 		category := r.URL.Path[len("/category/"):]
// 		username := sessions.GetCurrentUserName(r)
// 		context, err := db.GetTasks(username, "", category)
// 		categories := db.GetCategories(username)

// 		if err != nil {
// 			http.Redirect(w, r, "/", http.StatusInternalServerError)
// 		}
// 		if message != "" {
// 			context.Message = message
// 		}
// 		context.CSRFToken = "abcd"
// 		context.Categories = categories
// 		message = ""
// 		expiration := time.Now().Add(365 * 24 * time.Hour)
// 		cookie := http.Cookie{Name: "csrftoken", Value: "abcd", Expires: expiration}
// 		http.SetCookie(w, &cookie)
// 		homeTemplate.Execute(w, context)
// 	}
// }
