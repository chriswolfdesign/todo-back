package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	database "todo-back/db"

	"github.com/labstack/echo"
)

func GetHandler(ctx echo.Context, dm database.DatabaseManagerInterface) func(c echo.Context) error {
	return func(ctx echo.Context) error {
		param := ctx.Param("id")

		id, err := strconv.Atoi(param)
		if err != nil {
			ctxErr := ctx.String(http.StatusForbidden, "ID must be an integer")
			if ctxErr != nil {
				log.Println("unable to send context err:", ctxErr)
				return ctxErr
			}

			log.Printf("Could not parse ID %d: %s\n", id, err)
			return err
		}

		todo, err := dm.GetTodo(id)
		if err != nil {
			ctxErr := ctx.String(http.StatusForbidden, fmt.Sprintf("unable to find ID %d in database: %s\n", id, err))
			if ctxErr != nil {
				log.Println("unable to send context err:", ctxErr)
				return ctxErr
			}

			log.Printf("Could not parse ID %d: %s\n", id, err)
			return err

		}

		// If ID is 0, something went wrong
		if todo.ID == 0 {
			ctxErr := ctx.String(http.StatusForbidden, fmt.Sprintf("Could not find ID %d in database\n", id))
			if ctxErr != nil {
				log.Println("unable to send context err:", ctxErr)
				return ctxErr
			}

			log.Printf("unable to find ID %d in database\n", id)
			return nil
		}

		return ctx.JSON(http.StatusOK, todo)
	}
}
