package db

import (
	"database/sql"
	"fmt"

	model "todo-back/model"

	_ "github.com/lib/pq"
)

type DatabaseManager struct {
	*sql.DB
}

func CreateDatabase(host, port, user, password, dbName string) (*DatabaseManager, error) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		return nil, err
	}

	return &DatabaseManager{db}, nil
}

func (dm *DatabaseManager) GetList(table string) ([]model.TODOItem, error) {
	rows, err := dm.Query(fmt.Sprintf("select * from %s", table))
	if err != nil {
		return nil, err
	}

	list := []model.TODOItem{}

	for rows.Next() {
		var item model.TODOItem

		err := rows.Scan(&item.ID, &item.Body, &item.Completed)
		if err != nil {
			return nil, err
		}

		list = append(list, item)
	}

	return list, nil
}

func (dm *DatabaseManager) GetItem(table string, id int) (model.TODOItem, error) {
	rows, err := dm.Query(fmt.Sprintf("select * from %s where id=%d", table, id))
	if err != nil {
		return model.TODOItem{}, err
	}

	for rows.Next() {
		var item model.TODOItem

		err := rows.Scan(&item.ID, &item.Body, &item.Completed)
		if err != nil {
			return model.TODOItem{}, err
		}

		return item, nil
	}

	return model.TODOItem{}, nil
}

func (dm *DatabaseManager) UpdateCompletionStatus(table string, updateRequest model.UpdateRequest) error {
	sqlStatement := `UPDATE STREAM SET completed = $2 where id =$1`
	_, err := dm.Exec(sqlStatement, updateRequest.ID, updateRequest.Completed)
	return err
}

func (dm *DatabaseManager) CreateItem(table string, createRequest model.CreateRequest) error {
	sqlStatement := `INSERT INTO stream(body, completed) VALUES ($1, $2)`
	_, err := dm.Exec(sqlStatement, createRequest.Body, createRequest.Completed)
	return err
}

func (dm *DatabaseManager) DeleteItem(table string, id int) error {
	sqlStatement := `DELETE FROM stream WHERE id = $1`
	_, err := dm.Exec(sqlStatement, id)
	return err
}
