package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DataBase struct {
	database *sql.DB
}

func NewDataBase(dbUri string) (*DataBase, error) {
	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &DataBase{database: db}, nil
}

func (db *DataBase) Close() error {
	return db.database.Close()
}

func (db *DataBase) GetDataBase() *sql.DB {
	return db.database
}
