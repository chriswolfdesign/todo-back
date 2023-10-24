package database

import (
	"database/sql"
	"fmt"
	"log"
	"todo-back/config"
	"todo-back/model"

	_ "github.com/lib/pq"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . DatabaseManagerInterface
type DatabaseManagerInterface interface {
	EstablishDatabaseConnection(conf config.Config) error
	GetAllTodos() ([]model.Todo, error)
	GetTodo(id int) (*model.Todo, error)
	CreateTodo(text string, completed bool) (*model.Todo, error)
	CloseDatabase()
}

type DatabaseManager struct {
	db *sql.DB
}

func (dm *DatabaseManager) EstablishDatabaseConnection(conf config.Config) error {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.DBName)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Println("Could not connect to database:", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Could not ping database:", err)
		return err
	}

	log.Println("Connected!")
	dm.db = db

	return nil
}

func (dm *DatabaseManager) GetAllTodos() ([]model.Todo, error) {
	tx, err := dm.db.Begin()
	if err != nil {
		log.Println("unable to create transaction:", err)
		return nil, err
	}

	rows, err := tx.Query(`select * from todos`)
	if err != nil {
		log.Println("Unable to get all todos from database:", err)
		return nil, err
	}

	todos := []model.Todo{}

	for rows.Next() {
		todo := model.Todo{}

		err = rows.Scan(&todo.ID, &todo.Text, &todo.Completed)
		if err != nil {
			log.Println("Unable to get specific todo:", err)
			return nil, err
		}

		todos = append(todos, todo)
	}

	err = tx.Commit()
	if err != nil {
		log.Println("unable to commit transaction:", err)
		return nil, err
	}

	return todos, nil
}

func (dm *DatabaseManager) GetTodo(id int) (*model.Todo, error) {
	tx, err := dm.db.Begin()
	if err != nil {
		log.Println("unable to creae transaction in GetTodo handler:", err)
		return nil, err
	}

	query := `select * from todos where id=$1`
	rows, err := tx.Query(query, id)
	if err != nil {
		log.Printf("unable to get todo %d from database: %s\n", id, err)
		return nil, err
	}

	todo := model.Todo{}

	for rows.Next() {
		err = rows.Scan(&todo.ID, &todo.Text, &todo.Completed)
		if err != nil {
			log.Printf("unable to scan todo %d from database: %s\n", id, err)
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println("unable to commit transaction:", err)
		return nil, err
	}

	return &todo, nil
}

func (dm *DatabaseManager) CreateTodo(text string, completed bool) (*model.Todo, error) {
	tx, err := dm.db.Begin()
	if err != nil {
		log.Println("unable to create transaction in CreateTodo handler:", err)
		return nil, err
	}

	query := `insert into todos (text, completed) values ($1, $2) returning id, text, completed`
	todo := model.Todo{}

	err = tx.QueryRow(query, text, completed).Scan(&todo.ID, &todo.Text, &todo.Completed)
	if err != nil {
		log.Println("unable to scan when inserting into database:", err)
		return nil, err
	}

	return &todo, nil
}

func (dm *DatabaseManager) CloseDatabase() {
	dm.db.Close()
}
