package main

import (
	_ "github.com/byeol-i/battery-level-checker/docs" // echo-swagger middleware
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	//Swagger Ui middleware server
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1323"))
}