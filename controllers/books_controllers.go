package controllers

import (
	"books/conf"
	"books/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookReq struct {
	ID     int    `json:"id" param:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func GetBooks(c echo.Context) error {
	var books []models.Book

	db := conf.DBManager()
	db = db.Find(&books)

	return c.JSON(http.StatusOK, books)
}

func CreateBook(c echo.Context) (err error) {
	req := new(BookReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	book := models.Book{
		Title:  req.Title,
		Author: req.Author,
	}
	db := conf.DBManager()
	db = db.Create(&book)

	return c.JSON(http.StatusOK, book)
}

func GetBookByID(c echo.Context) (err error) {
	req := new(BookReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	var book []models.Book

	db := conf.DBManager()
	db = db.First(&book, req.ID)

	return c.JSON(http.StatusOK, book)

}

func UpdateBookByID(c echo.Context) (err error) {
	req := new(BookReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	book := models.Book{
		ID:     req.ID,
		Title:  req.Title,
		Author: req.Author,
	}

	db := conf.DBManager()
	db = db.Model(&book).Where("id = ?", req.ID).Updates(book)

	return c.JSON(http.StatusOK, book)

}

func DeleteBookByID(c echo.Context) (err error) {
	req := new(BookReq)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	var book []models.Book
	db := conf.DBManager()
	db = db.Delete(&book, req.ID)

	result := map[string]string{
		"respon_code": "200",
		"message":     "succsess",
	}
	return c.JSON(http.StatusOK, result)
}
