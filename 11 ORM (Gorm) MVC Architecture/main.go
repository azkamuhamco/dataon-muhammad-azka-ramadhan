package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "1234567",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

type User struct {
	ID       int    `json:"ID" form:"ID"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	gorm.Model
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

var users []User

// get all users
func GetUsersController(c echo.Context) error {
	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("ID"))

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get an user",
		"users":    users[ID],
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.ID = 0
	} else {
		newId := users[len(users)-1].ID + 1
		user.ID = newId
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	user.Name = name
	user.Email = email
	user.Password = password

	users = append(users, user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create a new user",
		"users":    users,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("id"))

	removeIndex := (func(s []User, index int) []User {
		ret := make([]User, 0)
		ret = append(ret, s[:index]...)
		return append(ret, s[index+1:]...)
	})

	users = removeIndex(users, ID)

	if err := DB.Save(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	ID, _ := strconv.Atoi(c.Param("ID"))

	users[ID].Name = c.FormValue("name")
	users[ID].Email = c.FormValue("email")
	users[ID].Password = c.FormValue("password")

	user.Name = users[ID].Name
	user.Email = users[ID].Email
	user.Password = users[ID].Password

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update an user",
		"user":     users[ID],
	})
}

func main() {
	// create a new echo instance
	e := echo.New()
	// route / to handler function
	e.GET("/users", GetUsersController)
	e.GET("/users/:ID", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("users/:ID", DeleteUserController)
	e.PUT("/users/:ID", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
