package main

import (
	"clean-azka/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	// create a new echo instance
	e := echo.New()
	// route / to handler function
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:ID", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("users/:ID", controller.DeleteUserController)
	e.PUT("/users/:ID", controller.UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
