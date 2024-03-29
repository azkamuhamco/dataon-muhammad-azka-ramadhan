package main

import (
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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

func ConnString() string {
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

	return connectionString
}

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", ConnString())
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
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

	if err := DB.Find(&users).Error; err != nil {
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
	usr := new(User)
	if err := c.Bind(usr); err != nil {
		return err
	}

	usr.Name = c.FormValue("name")
	usr.Email = c.FormValue("email")
	usr.Password = c.FormValue("password")

	sql := `INSERT INTO users(name, email, password, created_at, updated_at)
			VALUES("` + usr.Name + `", "` + usr.Email + `", "` + usr.Password + `",
				CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`

	DB.Exec(sql)

	usr.Name = c.FormValue("name")
	usr.Email = c.FormValue("email")
	usr.Password = c.FormValue("password")

	// users[len(users)-1].ID = uint(len(users)) + 1

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create an user",
		// "users":    users[len(users)],
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	requested_id := c.Param("ID")
	DB.Exec("DELETE FROM users WHERE ID = " + requested_id)

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "success delete ID "+string(requested_id))
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user := new(User)
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
