package main

import (
	"os"

	"dice.sawada.pro/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HTTPErrorHandler = handlers.CustomErrorHandler

	e.POST("/dice", handlers.Dice)

	port := os.Getenv("PORT")
	if port == "" {
		panic("port number is not define.")
	}

	e.Logger.Fatal(e.Start(":" + port))
}
