package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	db "todo-back/db"
	"todo-back/model"
)

func GenerateMuxServer(dm *db.DatabaseManager, table string) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	addGreetHandler(mux)
	addGetAllItemsHandler(mux, dm, table)
	addGetSingleTodoItemHandler(mux, dm, table)
	addCreateItemHandler(mux, dm, table)
	addDeleteItemHandler(mux, dm, table)

	return mux, nil
}

func addGreetHandler(mux *http.ServeMux) {
	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to our TODO List")
	})
}

func addGetAllItemsHandler(mux *http.ServeMux, dm *db.DatabaseManager, table string) {
	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Allow-Access-Control-Headers", "text/plain; application/json")

		list, err := dm.GetList(table)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(list)
	})
}

func addGetSingleTodoItemHandler(mux *http.ServeMux, dm *db.DatabaseManager, table string) {
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

			item, err := dm.GetItem(table, id)
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

			err = dm.UpdateCompletionStatus(table, updateRequest)
			if err != nil {
				log.Println("COULD NOT UPDATE COMPLETION STATUS:", err)
				w.WriteHeader(http.StatusBadRequest)
			}

			w.WriteHeader(http.StatusOK)
		}
	})
}

func addCreateItemHandler(mux *http.ServeMux, dm *db.DatabaseManager, table string) {
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

		err = dm.CreateItem(table, createRequest)
		if err != nil {
			log.Println("COULD NOT CREATE ITEM:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func addDeleteItemHandler(mux *http.ServeMux, dm *db.DatabaseManager, table string) {
	mux.HandleFunc("/deleteItem", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE")
		w.Header().Set("Allow-Access-Control-Headers", "text/plain; application/json")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal("COULD NOT GET BODY OF REQUEST:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var deleteRequest model.DeleteRequest
		json.Unmarshal(body, &deleteRequest)

		err = dm.DeleteItem(table, deleteRequest)
		if err != nil {
			log.Println("COULD NOT DELETE ITEM:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
