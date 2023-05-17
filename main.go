package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"go-web-project/database"
)

func main() {
	//e := echo.New()
	//
	//userGroup := e.Group("/user")

	dbURI := "postgres://postgres:qwerty54321@localhost:5432/go-web-project?sslmode=disable"

	db, err := database.NewDataBase(dbURI)
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		fmt.Errorf("cannot apply migrations: %w", err)
	}
	db.RunMigration()
}
