package routes

import (
	"server/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RideSubRoutes(g *echo.Group) {
	g.Use(middleware.JWT([]byte("secret")))
	g.POST("/", controllers.CreateRide)
	g.GET("/", controllers.GetRidesByUserID)
	g.GET("/:id", controllers.GetRideByRideID)
}
