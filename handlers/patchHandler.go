package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	database "todo-back/db"
	"todo-back/model"

	"github.com/labstack/echo"
)

func PatchHandler(ctx echo.Context, dm database.DatabaseManagerInterface) func(c echo.Context) error {
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

		postRequest := model.Request{}

		err = ctx.Bind(&postRequest)
		if err != nil {
			log.Println("unable to bind post request:", err)
			ctxErr := ctx.String(http.StatusForbidden, fmt.Sprintf("unable to bind patch request: %s\n", err))
			if ctxErr != nil {
				return ctxErr
			}

			return err
		}

		todo, err := dm.UpdateTodo(id, postRequest.Text, postRequest.Completed)
		if err != nil {
			log.Printf("unable to patch todo ID %d in database: %s\n", id, err)
			ctxErr := ctx.String(http.StatusBadRequest, fmt.Sprintf("unable to update todo ID %d in database: %s\n", id, err))
			if ctxErr != nil {
				return ctxErr
			}

			return err
		}

		err = ctx.JSON(http.StatusOK, todo)
		if err != nil {
			log.Println("unable to parse todo response:", err)
			ctxErr := ctx.String(http.StatusBadRequest, fmt.Sprintf("unable to parse todo response: %s\n", err))
			if ctxErr != nil {
				return ctxErr
			}

			return err
		}

		return nil
	}
}
