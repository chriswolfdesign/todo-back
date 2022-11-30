package main

import (
	"log"
	"net/http"
	"os"

	database "todo-back/db"
	endpoints "todo-back/endpoints"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	PSQL_PASSWORD := os.Getenv("PSQL_PASSWORD")
	PSQL_USER := os.Getenv("PSQL_USER")
	PSQL_PORT := os.Getenv("PSQL_PORT")
	PSQL_DB := os.Getenv("PSQL_DB")
	PSQL_TABLE := os.Getenv("PSQL_TABLE")
	PSQL_HOST := os.Getenv("PSQL_HOST")

	dm, err := database.CreateDatabase(PSQL_HOST, PSQL_PORT, PSQL_USER, PSQL_PASSWORD, PSQL_DB)
	defer dm.Close()

	if err != nil {
		log.Println("ERROR ACCESSING DATABASE")
		os.Exit(1)
	}

	log.Println("Contacted PSQL successfully")

	mux, err := endpoints.GenerateMuxServer(dm, PSQL_TABLE)
	if err != nil {
		log.Println("Could not generate mux server:", err)
		os.Exit(1)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedHeaders:     []string{"*"},
		AllowedMethods:     []string{"*"},
		AllowCredentials:   true,
		Debug:              true,
		OptionsPassthrough: true,
	})
	handler := c.Handler(mux)

	log.Println("ATTEMPTING TO START SERVER")

	err = http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal("ERROR STARTING SERVER:", err)
		os.Exit(1)
	}
}
