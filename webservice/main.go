package main
import (
	_ "fmt"
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
			AllowedMethods: []string{"GET", "POST", "OPTIONS"},
			OriginValidator: func(origin string, request *rest.Request) bool {
				return true
			},
			AllowedHeaders: []string{"Accept", "Content-Type", "X-Custom-Header", "Origin"},
			AccessControlAllowCredentials: true,
		})

		router, err := rest.MakeRouter(
			// -------
			rest.Get("/api/reminder", controller.GetAllReminders),
			rest.Get("/api/reminder/:id", controller.GetReminder),
			rest.Post("/api/reminder", controller.PostReminder),
			rest.Post("/api/del-reminder/:id", controller.DeleteReminder),
			rest.Post("/api/edit-reminder/:id", controller.PutReminder),
			// -------
		)

		if err != nil {
			log.Fatal(err)
		}
		api.SetApp(router)
		log.Println("running server on ", values.ServerPort)
		log.Fatal(http.ListenAndServe(values.ServerPort, api.MakeHandler()))

}
