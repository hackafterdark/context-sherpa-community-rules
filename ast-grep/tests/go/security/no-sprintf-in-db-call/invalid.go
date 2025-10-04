package main

import (
	"context"
	"database/sql"
	"fmt"
)

func vulnerableCode(db *sql.DB, petID int, pet Pet, categoryIDStr, userName, email string) error {
	ctx := context.Background()

	// This should be caught by our rule
	_, err := db.ExecContext(ctx, fmt.Sprintf(
		"INSERT INTO pets (id, name, category_id, status) VALUES (%d, '%s', %s, '%s')",
		petID, pet.Name, categoryIDStr, pet.Status,
	))
	if err != nil {
		return err
	}

	// This should also be caught
	rows, err := db.QueryContext(ctx, fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", userName))
	if err != nil {
		return err
	}
	defer rows.Close()

	// This should also be caught
	var id int
	err = db.QueryRowContext(ctx, fmt.Sprintf("SELECT id FROM users WHERE email = '%s'", email)).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}
