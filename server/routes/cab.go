package routes

import (
	"server/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func CabSubRoutes(g *echo.Group) {
	g.Use(middleware.JWT([]byte("secret")))
	g.POST("/", controllers.CreateCab)
	g.GET("/nearByCabs", controllers.FindNearByCab)
	g.GET("/:id", controllers.GetCabDetailsByID)
}
