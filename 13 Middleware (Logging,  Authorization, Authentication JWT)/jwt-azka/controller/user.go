package controller

import (
	"net/http"
	"strconv"
	"time"

	"jwt-azka/config"
	"jwt-azka/middlewares"
	"jwt-azka/model"

	"github.com/labstack/echo/v4"
)

var users []model.User
var guna = model.User{}
var DB = config.DB

func LoginUsers(user *model.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func LoginUsersController(c echo.Context) error {
	c.Bind(&guna)

	users, e := LoginUsers(&guna)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}

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
	guna.ID = uint(ID)

	if err := DB.First(&guna, ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get an user",
		"users":    &guna,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	usr := new(model.User)
	if err := c.Bind(usr); err != nil {
		return err
	}

	usr.Name = c.FormValue("name")
	usr.Email = c.FormValue("email")
	usr.Password = c.FormValue("password")

	sql := `INSERT INTO users(name, email, password, created_at, updated_at)
			VALUES("` + usr.Name + `", "` + usr.Email + `", "` + usr.Password + `",
				CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`

	if err := DB.Exec(sql).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create an user",
		"users":    usr,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	reqId, _ := strconv.Atoi(c.Param("ID"))

	if err := DB.Where("ID = ?", reqId).Delete(&guna).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "success delete ID "+strconv.Itoa(reqId))
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	ID, _ := strconv.Atoi(c.Param("ID"))

	user.ID = uint(ID)
	user.Name = c.FormValue("name")
	user.Email = c.FormValue("email")
	user.Password = c.FormValue("password")
	user.UpdatedAt = time.Now()

	doUpdate := DB.Table("users").Where("ID IN ?", []int{ID}).
		Updates(map[string]interface{}{
			"name":       user.Name,
			"email":      user.Email,
			"password":   user.Password,
			"updated_at": user.UpdatedAt,
		})

	if err := doUpdate.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update an user",
		"user":     user,
	})
}
