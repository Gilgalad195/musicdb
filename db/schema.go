package db

import (
	"database/sql"
	"log"
)

func InitDB(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS songs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		composer TEXT,
		first_line TEXT,
		themes TEXT,
		scripture_refs TEXT,
		pdf_path TEXT,
		media_path TEXT,
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}
