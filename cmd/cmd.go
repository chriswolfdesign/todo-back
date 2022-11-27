package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	database "todo-back/db"
	"todo-back/model"

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

	db, err := database.CreateDatabase(PSQL_HOST, PSQL_PORT, PSQL_USER, PSQL_PASSWORD, PSQL_DB)
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

		list, err := db.GetList(PSQL_TABLE)
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

		if r.Method == "GET" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET")
			w.Header().Set("Allow-Access-Control-Headers", "text/plain; application/json")

			id, err := strconv.Atoi(r.URL.Query().Get("id"))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}

			item, err := db.GetItem(PSQL_TABLE, id)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
		}

		if r.Method == "UPDATE" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "UPDATE")
			w.Header().Set("Allow-Access-Control-Headers", "text/plain; application/json")

			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Fatal("COULD NOT GET BODY OF REQUEST:", err)
				return
			}

			var updateRequest model.UpdateRequest
			json.Unmarshal(body, &updateRequest)

			err = db.UpdateCompletionStatus(PSQL_TABLE, updateRequest)
			if err != nil {
				log.Println("COULD NOT UPDATE COMPLETION STATUS:", err)
				w.WriteHeader(http.StatusBadRequest)
			}

			w.WriteHeader(http.StatusOK)
		}
	})

	mux.HandleFunc("/createItem", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "PUT")
		w.Header().Set("Allow-Access-Control-Headers", "text/plain; application/json")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal("COULD NOT GET BODY OF REQUEST:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var createRequest model.CreateRequest
		json.Unmarshal(body, &createRequest)

		err = db.CreateItem(PSQL_TABLE, createRequest)
		if err != nil {
			log.Println("COULD NOTE CREATE ITEM:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/deleteItem", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "PUT")
		w.Header().Set("Allow-Access-Control-Headers", "text/plain; application/json")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal("COULD NOT GET BODY OF REQUEST:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var deleteRequest model.DeleteRequest
		json.Unmarshal(body, &deleteRequest)

		err = db.DeleteItem(PSQL_TABLE, deleteRequest)
		if err != nil {
			log.Println("COULD NOT DELETE ITEM:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
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