package main

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------- controller --------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get an user",
		"user":     users[id],
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	removeIndex := (func(s []User, index int) []User {
		ret := make([]User, 0)
		ret = append(ret, s[:index]...)
		return append(ret, s[index+1:]...)
	})

	users = removeIndex(users, id)
	return c.NoContent(http.StatusNoContent)
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))

	users[id].Name = c.FormValue("name")
	users[id].Email = c.FormValue("email")
	users[id].Password = c.FormValue("password")

	user.Name = users[id].Name
	user.Email = users[id].Email
	user.Password = users[id].Password

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update an user",
		"user":     users[id],
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 0
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	user.Name = name
	user.Email = email
	user.Password = password

	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create a new user",
		"user":     user,
	})
}

func main() {
	e := echo.New()
	// routing with query param
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
