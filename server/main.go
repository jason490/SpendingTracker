package main

import (
	"SpendingTracker/internal/frontend"
	"SpendingTracker/internal/server"
	"SpendingTracker/internal/storage"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Opens the init.sql file
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
	fmt.Println("Clearing Database ... ")

	store := storage.NewSqliteStorage(db)

	// Run the api server
	e := server.RunServer(store)
	frontend.RunFrontend(store, e)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		e.Logger.Fatal(e.Start(":8080"))
	}()
	<-ctx.Done()
}
