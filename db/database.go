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
	rows, err := dm.db.Query(`select * from todos`)
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

	return todos, nil
}

func (dm *DatabaseManager) CloseDatabase() {
	dm.db.Close()
}
