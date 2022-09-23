package main

import (
	"api/apis"
	"api/controller"
	"fmt"

	"github.com/jasonlvhit/gocron"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	gocron.Every(3).Seconds().Do(startLastNews)
	gocron.Every(10).Seconds().Do(apis.CreateNews)
	<-gocron.Start()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/api/get/news/:id", controller.GetNewsByID)
	e.GET("/api/get/all/news", controller.GetAllNews)
	e.Logger.Fatal(e.Start(":8081"))
}

func startLastNews() {
	lastNews, err := apis.GetLastNews()
	fmt.Println(lastNews)
	apis.HandleError(err, "Line 123")
	apis.SendLastNews(lastNews)
}
