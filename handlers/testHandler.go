package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func GetTest(c echo.Context) error {
	log.Println("/test GET endpoint has been reached")
	return c.String(http.StatusOK, "Echo server running correctly")
}
