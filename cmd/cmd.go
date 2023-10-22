package main

import (
	"todo-back/config"
	database "todo-back/db"
	"todo-back/handlers"

	"github.com/labstack/echo"
)

const (
	CONFIG_PATH = "config/config.yaml"
)

func main() {
	e := echo.New()

	conf := config.ReadConfigFile(CONFIG_PATH)

	database.GenerateDatabaseConnection(conf)

	e.GET("/test", handlers.GetTest)
	e.Logger.Fatal(e.Start(":3000"))
}
