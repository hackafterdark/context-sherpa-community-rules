package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func badQuery(db *sql.DB, userID int) {
	// This is vulnerable - uses string formatting in query
	row := db.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE id = %d", userID))
	var name string
	err := row.Scan(&name)
	if err != nil {
		log.Printf("Error scanning row: %v", err)
		return
	}
}

func badExec(db *sql.DB, email string) error {
	// This is vulnerable - uses string formatting in exec
	_, err := db.Exec(fmt.Sprintf("INSERT INTO users (email) VALUES ('%s')", email))
	return err
}

func badExecContext(ctx context.Context, db *sql.DB, name string) error {
	// This is vulnerable - uses string formatting in exec context
	_, err := db.ExecContext(ctx, fmt.Sprintf("UPDATE users SET name = '%s' WHERE id = 1", name))
	return err
}

func badQueryRow(db *sql.DB, tableName string) {
	// This is vulnerable - uses string formatting in query row
	row := db.QueryRow(fmt.Sprintf("SELECT * FROM %s", tableName))
	var count int
	row.Scan(&count)
}

func badQueryRowContext(ctx context.Context, db *sql.DB, userInput string) {
	// This is vulnerable - uses string formatting in query row context
	row := db.QueryRowContext(ctx, fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", userInput))
	var id int
	row.Scan(&id)
}
