package main

import (
	"flag"
	"log"
	"strings"

	"./config"
	controller "./controller"
	model "./model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	//read config http
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
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	//api need authorization
	r := e.Group("/api")
	r.Use(middleware.JWTWithConfig(config))

	// admin management
	r.GET("/administrators", controller.GetAllAdmin)
	r.GET("/administrators/:id", controller.GetAdminById)
	r.POST("/administrators", controller.PostAdmin)
	r.POST("/del-administrators/:id", controller.DeleteAdmin)
	r.POST("/edit-administrators/:id", controller.PutAdmin)
	// ACL management
	r.GET("/accessrules", controller.GetAllAccessRules)
	r.GET("/search-acl", controller.SearchACLCtrl)
	r.GET("/accessrules/:id", controller.GetAccessRuleById)
	r.POST("/accessrules", controller.PostAccessRule)
	r.POST("/del-accessrules/:id", controller.DeleteAccessRule)
	r.POST("/edit-accessrules/:id", controller.PutAccessRule)
	// Accepting Host management
	r.GET("/acceptinghosts", controller.GetAcceptingHost)
	r.GET("/acceptinghosts/:id", controller.GetAcceptingHostById)
	r.POST("/acceptinghosts", controller.PostAcceptingHost)
	r.POST("/del-acceptinghosts/:id", controller.DeleteAcceptingHost)
	r.POST("/edit-acceptinghosts/:id", controller.PutAcceptingHost)
	// Applications management
	r.GET("/applications", controller.GetAllApplications)
	r.GET("/applications/:id", controller.GetApplicationsById)
	r.POST("/applications", controller.PostApplications)
	r.POST("/del-applications/:id", controller.DeleteApplications)
	r.POST("/edit-applications/:id", controller.PutApplications)
	// Start server
	e.Logger.Fatal(e.Start(values.ServerPort))
}
