package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

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

	fmt.Println("PASSWORD:", PSQL_PASSWORD)
	fmt.Println("USER:", PSQL_USER)
	fmt.Println("PORT:", PSQL_PORT)
	fmt.Println("DB:", PSQL_DB)
	fmt.Println("TABLE:", PSQL_TABLE)
	fmt.Println("HOST:", PSQL_HOST)

	db, err := createDatabase(PSQL_HOST, PSQL_PORT, PSQL_USER, PSQL_PASSWORD, PSQL_DB)
	defer db.Close()

	if err != nil {
		log.Println("ERROR ACCESSING DATABASE")
		os.Exit(1)
	}

	log.Println("Contacted PSQL successfully")

	mux := http.NewServeMux()

	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to our TODO List")
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://foo.com", "http://foo.com:8080"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := c.Handler(mux)

	log.Println("ATTEMPTING TO START SERVER")

	err = http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal("ERROR STARTING SERVER:", err)
		os.Exit(1)
	}
}

func createDatabase(host, port, user, password, dbName string) (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	return sql.Open("postgres", psqlconn)
}
