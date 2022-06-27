package controller

import (
	"net/http"
	"strconv"
	"time"

	"jwt-azka/model"

	"gorm.io/datatypes"

	"github.com/labstack/echo/v4"
)

var books []model.Book
var buku = model.Book{}
var bok = new(model.Book)

// get all users
func GetBooks(c echo.Context) error {
	if err := DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all books",
		"users":    books,
	})
}

// get user by id
func GetBook(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("ID"))
	buku.ID = uint(ID)

	if err := DB.First(&buku, ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get an book",
		"users":    &buku,
	})
}

// create new user
func CreateBook(c echo.Context) error {
	if err := c.Bind(bok); err != nil {
		return err
	}

	intAuthorID, _ := strconv.Atoi(c.FormValue("author_id"))
	intPublisherID, _ := strconv.Atoi(c.FormValue("publisher_id"))
	timePublishedAt, _ := time.Parse("2006-01-02", c.FormValue("published_at"))

	bok.Title = c.FormValue("title")
	bok.AuthorID = uint(intAuthorID)
	bok.PublisherID = uint(intPublisherID)
	bok.Price, _ = strconv.Atoi(c.FormValue("price"))
	bok.Page, _ = strconv.Atoi(c.FormValue("page"))
	bok.Weight, _ = strconv.Atoi(c.FormValue("weight"))
	bok.PublishedAt = datatypes.Date(timePublishedAt)

	strAuthorID := strconv.Itoa(int(bok.AuthorID))
	strPublisherID := strconv.Itoa(int(bok.PublisherID))
	strPrice := strconv.Itoa(bok.Price)
	strPage := strconv.Itoa(bok.Page)
	strWeight := strconv.Itoa(bok.Weight)

	sql := `INSERT INTO books(title, author_id, publisher_id, price, page, weight, published_at, created_at,
			updated_at) VALUES("` + bok.Title + `", ` + strAuthorID + `, ` + strPublisherID + `, ` + strPrice + `,
			` + strPage + `,` + strWeight + `, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`

	if err := DB.Exec(sql).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create an book",
		"users":    bok,
	})
}

// delete user by id
func DeleteBook(c echo.Context) error {
	reqId, _ := strconv.Atoi(c.Param("ID"))

	if err := DB.Where("ID = ?", reqId).Delete(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "success delete ID book "+strconv.Itoa(reqId))
}

// update user by id
func UpdateBook(c echo.Context) error {
	if err := c.Bind(bok); err != nil {
		return err
	}

	ID, _ := strconv.Atoi(c.Param("ID"))
	intAuthorID, _ := strconv.Atoi(c.FormValue("author_id"))
	intPublisherID, _ := strconv.Atoi(c.FormValue("publisher_id"))
	timePublishedAt, _ := time.Parse("2006-01-02", c.FormValue("published_at"))

	bok.ID = uint(ID)
	bok.Title = c.FormValue("title")
	bok.AuthorID = uint(intAuthorID)
	bok.PublisherID = uint(intPublisherID)
	bok.Price, _ = strconv.Atoi(c.FormValue("price"))
	bok.Page, _ = strconv.Atoi(c.FormValue("page"))
	bok.Weight, _ = strconv.Atoi(c.FormValue("weight"))
	bok.PublishedAt = datatypes.Date(timePublishedAt)
	bok.UpdatedAt = time.Now()

	doUpdate := DB.Table("books").Where("ID IN ?", []int{ID}).
		Updates(map[string]interface{}{
			"title":        bok.Title,
			"author_id":    bok.AuthorID,
			"publisher_id": bok.PublisherID,
			"price":        bok.Price,
			"page":         bok.Page,
			"weight":       bok.Weight,
			"published_at": bok.PublishedAt,
			"updated_at":   bok.UpdatedAt,
		})

	if err := doUpdate.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update an book",
		"user":     bok,
	})
}
