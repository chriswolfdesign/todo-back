package handlers

import (
	"fmt"
	"log"
	"net/http"
	database "todo-back/db"
	"todo-back/model"

	"github.com/labstack/echo"
)

func CreateHandler(ctx echo.Context, dm database.DatabaseManagerInterface) func(c echo.Context) error {
	return func(ctx echo.Context) error {
		postRequest := model.Request{}

		err := ctx.Bind(&postRequest)
		if err != nil {
			log.Println("unable to bind post request:", err)
			ctxErr := ctx.String(http.StatusForbidden, fmt.Sprintf("unable to bind post request: %s\n", err))
			if ctxErr != nil {
				return ctxErr
			}

			return err
		}

		todo, err := dm.CreateTodo(postRequest.Text, postRequest.Completed)
		if err != nil {
			log.Println("unable to create todo in database:", err)
			ctxErr := ctx.String(http.StatusBadRequest, fmt.Sprintf("unable to create todo in database: %s\n", err))
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
