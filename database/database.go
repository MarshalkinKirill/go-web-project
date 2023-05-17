package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

func (db *DataBase) RunMigration() error {
	driver, err := postgres.WithInstance(db.database, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create database driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+"/migrations",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
	//migrationsDir := "C:\\Users\\evilr\\GolandProjects\\go-web-project\\database\\migrations"
	//
	//// Применение миграций с использованием goose
	//err := goose.Up(db.database, migrationsDir)
	//if err != nil {
	//	return err
	//}
	//
	//return nil
}
