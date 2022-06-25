package routes

import (
	c "jwt-azka/controller"
	m "jwt-azka/middlewares"

	"github.com/labstack/echo/v4"
	echoMid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	e.POST("/users", c.CreateUserController)

	// implement middleware with group routing
	eAuth := e.Group("")
	eAuth.Use(echoMid.BasicAuth(m.BasicAuthDB))

	// route / to handler function
	eAuth.GET("/users", c.GetUsersController)
	eAuth.GET("/users/:ID", c.GetUserController)
	eAuth.DELETE("users/:ID", c.DeleteUserController)
	eAuth.PUT("/users/:ID", c.UpdateUserController)

	return e
}
