package routes

import (
	"server/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func UserSubRoutes(g *echo.Group) {
	// g.Use(helpers.JWTAuthentication)

	g.Use(middleware.JWT([]byte("secret")))

	// g.GET("/", controllers.GetUsers)
	g.GET("", controllers.GetUserByID)
	// g.POST("/login", controllers.Login)

	g.PUT("/", controllers.UpdateUserById)

}
