package main

import (
	"flag"
	"log"
	"strings"
	"./config"
	controller "./controller"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
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
  // Echo instance
  e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
     AllowOrigins: []string{"*"},
     AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
 }))

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
	e.POST("/api/login", controller.LoginCtrl)
	e.POST("/api/search", controller.SearchCtrl)
	// admin management
  e.GET("/api/administrators", controller.GetAllAdmin)
  e.GET("/api/administrators/:id", controller.GetAdminById)
  e.POST("/api/administrators", controller.PostAdmin)
  e.POST("/api/edit-administrators/:id", controller.PutAdmin)
  // Start server
  e.Logger.Fatal(e.Start(values.ServerPort))
}
