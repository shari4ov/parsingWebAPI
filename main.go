package main

import (
	"api/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/api/get/news/:id", controller.GetNewsByID)
	e.GET("/api/get/all/news", controller.GetAllNews)
	e.Logger.Fatal(e.Start(":8081"))
}
