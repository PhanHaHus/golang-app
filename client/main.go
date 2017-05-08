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
	// http.HandleFunc("/login", views.LoginFunc)
	// http.HandleFunc("/logout", views.RequiresLogin(views.LogoutFunc))
	//these handlers fetch set of tasks
	http.HandleFunc("/",(views.ShowAllTasksFunc))
	// http.HandleFunc("/add-admin", views.RequiresLogin(views.AddReminder))
	log.Println("running server on ", values.ServerPort)
	log.Fatal(http.ListenAndServe(values.ServerPort, nil))
}
