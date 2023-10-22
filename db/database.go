package database

import (
	"database/sql"
	"fmt"
	"log"
	"todo-back/config"

	_ "github.com/lib/pq"
)

type DatabaseManagerInterface interface {
	EstablishDatabaseConnection(conf config.Config) error
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
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("Could not ping database:", err)
		return err
	}

	log.Println("Connected!")
	dm.db = db

	return nil
}
