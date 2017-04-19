package main
import (
	"flag"
	"log"
	"net/http"
	"strings"
	config "./config"
	views "./views"
)

func main() {
	values, err := config.ReadConfig("config.json")
	var port *string
	if err != nil {
		port = flag.String("port", "", "IP address")
		flag.Parse()
		//User is expected to give :8080 like input, if they give 8080
		//we'll append the required ':'
		if !strings.HasPrefix(*port, ":") {
			*port = ":" + *port
			log.Println("port is " + *port)
		}
		values.ServerPort = *port
	}
	http.Handle("/static/", http.FileServer(http.Dir("public")))
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))

	views.PopulateTemplates()

	//Login logout
	http.HandleFunc("/login", views.LoginFunc)
	http.HandleFunc("/logout", views.RequiresLogin(views.LogoutFunc))
	// http.HandleFunc("/signup/", views.SignUpFunc)
	//these handlers fetch set of tasks
	http.HandleFunc("/",views.RequiresLogin(views.ShowAllTasksFunc))
	http.HandleFunc("/add-admin", views.RequiresLogin(views.AddReminder))
	http.HandleFunc("/detail-admin/", views.RequiresLogin(views.DetailReminderFunc))
	http.HandleFunc("/edit-admin/", views.RequiresLogin(views.EditReminderFunc))
	// //these handlers perform action like delete, mark as complete etc

	// http.HandleFunc("/search/", views.RequiresLogin(views.SearchTaskFunc))
	log.Println("running server on ", values.ServerPort)
	log.Fatal(http.ListenAndServe(values.ServerPort, nil))
}
