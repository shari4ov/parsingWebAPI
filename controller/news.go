package controller

import (
	api "api/apis"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetNewsByID(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(c)
	return c.JSON(http.StatusOK, api.GetNewsById(id))
}
func GetAllNews(c echo.Context) error {
	fmt.Println(c)
	return c.JSON(http.StatusOK, api.GetAllNews())
}
