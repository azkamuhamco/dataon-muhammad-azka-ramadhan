package routes

import (
	"jwt-azka/constants"
	c "jwt-azka/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// echoMid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// // implement middleware with group routing
	// eAuth := e.Group("")
	// eAuth.Use(echoMid.BasicAuth(m.BasicAuthDB))

	// JWT Group
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// route users
	r.GET("/users", c.GetUsersController)
	r.GET("/users/:ID", c.GetUserController)
	e.POST("/users", c.CreateUserController)
	r.DELETE("users/:ID", c.DeleteUserController)
	r.PUT("/users/:ID", c.UpdateUserController)

	// route books
	e.GET("/books", c.GetBooks)
	e.GET("/books/:ID", c.GetBook)
	r.POST("/books", c.CreateBook)
	r.DELETE("books/:ID", c.DeleteBook)
	r.PUT("/books/:ID", c.UpdateBook)

	return e
}
