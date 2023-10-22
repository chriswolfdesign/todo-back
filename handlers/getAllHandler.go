package handlers

import (
	"log"
	"net/http"
	database "todo-back/db"
	"todo-back/model"

	"github.com/labstack/echo"
)

func GetAllHandler(ctx echo.Context, dm database.DatabaseManagerInterface) func(c echo.Context) error {
	return func(ctx echo.Context) error {
		todos, err := dm.GetAllTodos()
		if err != nil {
			log.Println("Unable to get all todos from handler:", err)
			ctx.String(http.StatusInternalServerError, "Unable to retrieve list of todo items")
			return err
		}

		payload := map[string][]model.Todo{
			"todos": todos,
		}

		log.Println("Sending all todos via GetAllHandler")
		return ctx.JSON(http.StatusOK, payload)
	}
}
