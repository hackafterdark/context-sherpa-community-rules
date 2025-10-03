package main

import (
	"context"
	"database/sql"
	"log"
)

func goodQuery(db *sql.DB, userID string) {
	// This is safe - uses parameterized query
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", userID)
	var name string
	err := row.Scan(&name)
	if err != nil {
		log.Printf("Error scanning row: %v", err)
		return
	}
}

func goodExec(db *sql.DB, email string) error {
	// This is safe - uses parameterized query
	_, err := db.Exec("INSERT INTO users (email) VALUES (?)", email)
	return err
}

func goodExecContext(ctx context.Context, db *sql.DB, name string) error {
	// This is safe - uses parameterized query with context
	_, err := db.ExecContext(ctx, "UPDATE users SET name = ? WHERE id = 1", name)
	return err
}
