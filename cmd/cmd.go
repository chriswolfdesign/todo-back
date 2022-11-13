package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type TODOItem struct {
	ID        string `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

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

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Allow-Access-Control-Headers", "text/plain; application/json")

		list, err := getList(db, PSQL_TABLE)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(list)
	})

	mux.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Allow-Access-Control-Headers", "text/plain; application/json")

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		item, err := getItem(db, PSQL_TABLE, id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
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

func getList(db *sql.DB, table string) ([]TODOItem, error) {
	rows, err := db.Query(fmt.Sprintf("select * from %s", table))
	if err != nil {
		return nil, err
	}

	list := []TODOItem{}

	for rows.Next() {
		var item TODOItem

		err := rows.Scan(&item.ID, &item.Body, &item.Completed)
		if err != nil {
			return nil, err
		}

		list = append(list, item)
	}

	return list, nil
}

func getItem(db *sql.DB, table string, id int) (TODOItem, error) {
	rows, err := db.Query(fmt.Sprintf("select * from %s where id=%d", table, id))
	if err != nil {
		return TODOItem{}, err
	}

	for rows.Next() {
		var item TODOItem

		err := rows.Scan(&item.ID, &item.Body, &item.Completed)
		if err != nil {
			return TODOItem{}, err
		}

		return item, nil
	}

	return TODOItem{}, nil
}
