package route

import (
	"books/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/books", controllers.GetBooks)
	e.POST("/books", controllers.CreateBook)
	e.GET("/books/:id", controllers.GetBookByID)
	e.PUT("/books/:id", controllers.UpdateBookByID)
	e.DELETE("/books/:id", controllers.DeleteBookByID)
	return e
}
