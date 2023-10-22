package main

import (
	"log"
	"os"
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

	var dm database.DatabaseManagerInterface

	dm = &database.DatabaseManager{}

	err := dm.EstablishDatabaseConnection(conf)
	if err != nil {
		log.Fatal("Unable to establish database connection:", err)
		os.Exit(1)
	}

	e.GET("/test", handlers.GetTest)
	e.Logger.Fatal(e.Start(":3000"))
}
