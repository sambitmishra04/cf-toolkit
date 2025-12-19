package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"golang.org/x/oauth2"
)

var db *sql.DB

func initDB() {
	var connStr string
	if url := os.Getenv("DATABASE_URL"); url != "" {
		connStr = url
	} else {
		host := "localhost"
		if os.Getenv("DB_HOST") != "" {
			host = os.Getenv("DB_HOST")
		}
		connStr = fmt.Sprintf("postgres://cfuser:cfpassword@%s:5432/cf_toolkit?sslmode=disable", host)
	}

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
    );
    
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) UNIQUE NOT NULL,
        access_token TEXT NOT NULL,
        refresh_token TEXT NOT NULL,
        token_expiry TIMESTAMP NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
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

func saveUser(email string, token *oauth2.Token) error {
	query := `
		INSERT INTO users (email, access_token, refresh_token, token_expiry)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (email) DO UPDATE SET
			access_token = EXCLUDED.access_token,
			refresh_token = EXCLUDED.refresh_token,
			token_expiry = EXCLUDED.token_expiry;
	`
	_, err := db.Exec(query, email, token.AccessToken, token.RefreshToken, token.Expiry)
	return err
}

type User struct {
	Email        string
	AccessToken  string
	RefreshToken string
	TokenExpiry  time.Time
}

func getAllUsers() ([]User, error) {
	rows, err := db.Query("SELECT email, access_token, refresh_token, token_expiry FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Email, &u.AccessToken, &u.RefreshToken, &u.TokenExpiry); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
