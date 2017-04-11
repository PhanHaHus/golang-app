package main
import (
	"github.com/ant0ine/go-json-rest/rest"
	"flag"
	"log"
	"net/http"
	"strings"
	"./config"
	controller "./controller"
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

	api := rest.NewApi()
		api.Use(rest.DefaultDevStack...)
		api.Use(&rest.CorsMiddleware{ //CorsMiddleware
			RejectNonCorsRequests: false,
			AllowedMethods: []string{"GET", "POST", "PUT"},
			AllowedHeaders: []string{
				"Accept", "Content-Type", "X-Custom-Header", "Origin"},
			AccessControlAllowCredentials: true,
			AccessControlMaxAge:           3600,
		})

		router, err := rest.MakeRouter(
			rest.Get("/api/get-task", controller.GetAllReminders),
			// rest.Post("/api/add-task", controller.AddTaskFuncAPI),
			// rest.Delete("/api/get-deleted-task", controller.GetDeletedTaskFuncAPI),
			// rest.Put("/api/update-task", controller.UpdateTaskFuncAPI),
			// rest.Delete("/reminders/:id", i.DeleteReminder),
		)

		if err != nil {
			log.Fatal(err)
		}
		api.SetApp(router)
		log.Println("running server on ", values.ServerPort)
		log.Fatal(http.ListenAndServe(values.ServerPort, api.MakeHandler()))

}
