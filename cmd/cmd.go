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

	conf, err := config.ReadConfigFile(CONFIG_PATH)
	if err != nil {
		log.Println("unable to open config")
		os.Exit(1)
	}

	dm := &database.DatabaseManager{}

	err = dm.EstablishDatabaseConnection(*conf)
	if err != nil {
		log.Fatal("Unable to establish database connection:", err)
		os.Exit(1)
	}
	defer dm.CloseDatabase()

	e.GET("/test", handlers.GetTest)
	e.GET("/todos", handlers.GetAllHandler(e.AcquireContext(), dm))
	e.GET("/todo/:id", handlers.GetHandler(e.AcquireContext(), dm))
	e.POST("/todo", handlers.CreateHandler(e.AcquireContext(), dm))
	e.DELETE("/todo/:id", handlers.DeleteHandler(e.AcquireContext(), dm))
	e.PATCH("/todo/:id", handlers.PatchHandler(e.AcquireContext(), dm))
	e.Logger.Fatal(e.Start(":3000"))
}
