package main

import (
	"fmt"
	"server/controllers"
	"server/routes"

	"github.com/Kamva/mgm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	routes.UserSubRoutes(e.Group("/users"))
	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)

	e.Logger.Fatal(e.Start(":4000"))
}

func init() {
	err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI("mongodb://localhost:27017/stack"))
	if err != nil {
		fmt.Println("connected")
	}
}
