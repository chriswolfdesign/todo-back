package main

import (
	"todo-back/handlers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/test", handlers.GetTest)
	e.Logger.Fatal(e.Start(":3000"))
}
