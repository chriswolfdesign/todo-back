package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"todo-back/config"

	_ "github.com/lib/pq"
)

func GenerateDatabaseConnection(conf config.Config) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.DBName)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Println("Could not connect to database:", err)
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("Could not ping database:", err)
		os.Exit(1)
	}

	log.Println("Connected!")
}
