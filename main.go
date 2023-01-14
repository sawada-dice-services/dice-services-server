package main

import (
	"fmt"
	"os"

	"dice.sawada.pro/handlers"
	"github.com/labstack/echo/v4"
)

var DEFAULT_PORT = "8080"

func main() {
	e := echo.New()
	e.HTTPErrorHandler = handlers.CustomErrorHandler

	e.POST("/dice", handlers.Dice)

	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
		fmt.Printf("WARN: port number is not defined. listening %s", DEFAULT_PORT)
	}

	e.Logger.Fatal(e.Start(":" + port))
}
