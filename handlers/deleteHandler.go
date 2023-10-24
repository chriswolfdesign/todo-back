package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	database "todo-back/db"

	"github.com/labstack/echo"
)

func DeleteHandler(ctx echo.Context, dm database.DatabaseManagerInterface) func(e echo.Context) error {
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

		err = dm.DeleteTodo(id)

		if err != nil {
			ctxErr := ctx.String(http.StatusForbidden, fmt.Sprintf("unable to delete ID %d from database: %s\n", id, err))
			if ctxErr != nil {
				log.Println("unable to send context err:", ctxErr)
				return ctxErr
			}

			log.Printf("Could not parse ID %d: %s\n", id, err)
			return err
		}

		ctxErr := ctx.String(http.StatusOK, fmt.Sprintf("successfully deleted ID %d from database\n", id))
		if ctxErr != nil {
			log.Println("unable to send context err:", ctxErr)
			return ctxErr
		}

		return nil
	}
}
