package main

import (
	"fmt"
	"os"
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
}
