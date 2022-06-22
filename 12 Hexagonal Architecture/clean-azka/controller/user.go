package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"clean-azka/model"

	"github.com/labstack/echo/v4"
)

var users []model.User
var debe = model.DB

// get all users
func GetUsersController(c echo.Context) error {
	if err := debe.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for i := 0; i < len(users); i++ {
		users[i].ID = uint(i + 1)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("ID"))
	ID -= 1

	if err := debe.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users[ID].ID = uint(ID) + 1

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get an user",
		"users":    users[ID],
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	usr := new(model.User)
	if err := c.Bind(usr); err != nil {
		return err
	}

	DBS, er := sql.Open("mysql", model.ConnString())
	if er != nil {
		panic(er)
	}
	defer DBS.Close()

	sql := "INSERT INTO users(name, email, password, created_at, updated_at) VALUES(?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	stmt, err := DBS.Prepare(sql)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()

	usr.Name = c.FormValue("name")
	usr.Email = c.FormValue("email")
	usr.Password = c.FormValue("password")

	result, err2 := stmt.Exec(usr.Name, usr.Email, usr.Password)

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	users[len(users)-1].ID = uint(len(users)) + 1

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get an user",
		"users":    users[len(users)-1],
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	DBS, er := sql.Open("mysql", model.ConnString())
	if er != nil {
		panic(er)
	}
	defer DBS.Close()

	requested_id := c.Param("ID")
	sql := "DELETE FROM users WHERE ID = ?"
	stmt, err := DBS.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}

	_, err2 := stmt.Exec(requested_id)
	if err2 != nil {
		panic(err2)
	}

	return c.JSON(http.StatusOK, "success delete ID "+string(requested_id))
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	ID, _ := strconv.Atoi(c.Param("ID"))
	ID -= 1

	users[ID].Name = c.FormValue("name")
	users[ID].Email = c.FormValue("email")
	users[ID].Password = c.FormValue("password")

	user.ID = uint(ID) + 1
	user.Name = users[ID].Name
	user.Email = users[ID].Email
	user.Password = users[ID].Password

	if err := debe.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update an user",
		"user":     users[ID],
	})
}
