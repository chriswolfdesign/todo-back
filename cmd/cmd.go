package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/test", getTest)
	e.Logger.Fatal(e.Start(":3000"))
}

func getTest(c echo.Context) error {
	log.Println("/test GET endpoint has been reached")
	c.String(http.StatusOK, "Echo server running correctly")
	return nil
}
