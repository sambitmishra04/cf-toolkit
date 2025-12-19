package main

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func initDB() {
	host := "localhost"
	if os.Getenv("DB_HOST") != "" {
		host = os.Getenv("DB_HOST")
	}

	connStr := fmt.Sprintf("postgres://cfuser:cfpassword@%s:5432/cf_toolkit?sslmode=disable", host)
	
	var err error
	db, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Database unreachable: %v", err)
	}

	createTable()
}

func createTable() {
	query := `
    CREATE TABLE IF NOT EXISTS contest_events (
        contest_id INTEGER PRIMARY KEY,
        added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func eventExists(contestID int) bool {
	var exists bool
    query := "SELECT EXISTS(SELECT 1 FROM contest_events WHERE contest_id = $1)"
    err := db.QueryRow(query, contestID).Scan(&exists)
    if err != nil {
        log.Printf("Error checking DB: %v", err)
        return false
    }
    return exists
}

func saveEvent(contestID int) {
	query := "INSERT INTO contest_events (contest_id) VALUES ($1)"
    _, err := db.Exec(query, contestID)
    if err != nil {
        log.Printf("Error saving event: %v", err)
    }
}
