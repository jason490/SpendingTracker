package server

import (
	"database/sql"

	"log"
	"os"
	"SpendingTracker/internal/storage"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	store *storage.Storage
	e     *echo.Echo
}

func RunServer() *echo.Echo {
	// Add condition if there is no db file
	initFile, err := os.ReadFile("./init.sql")
	if err != nil {
		log.Fatal("Missing init.sql file!")
	}

	db, err := sql.Open("sqlite3", "./database/main.db")
	if err != nil {
		log.Fatal("Unable to open main.db!")
	}
	defer db.Close()
	// Add a condition that doesn't clear the db
	// This clears the db
	if _, err := db.Exec(string(initFile)); err != nil {
		log.Fatal("Unable to execute init sql")
	}

	store := storage.NewSqliteStorage(db)

	e := echo.New()
	server := Server{
		store: store,
		e:     e,
	}
	server.routes()
	return e
}
