package main

import (
	"context"
	"database/sql"
)

func safeCode(db *sql.DB, petID int, pet Pet) error {
	ctx := context.Background()

	// This should NOT be caught (safe parameterized query)
	_, err := db.ExecContext(ctx, "INSERT INTO pets (id, name) VALUES ($1, $2)", petID, pet.Name)
	return err
}
