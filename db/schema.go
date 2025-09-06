package db

import (
	"database/sql"
	"log"
)

const createSongsTableSQL = `
	CREATE TABLE IF NOT EXISTS songs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		composer TEXT,
		first_line TEXT,
		themes TEXT,
		scripture_refs TEXT,
		pdf_path TEXT,
		lyric_sheet_path TEXT,
		media_path TEXT,
		delete_date TEXT
	);
`

const createMetaTableSQL = `
	CREATE TABLE IF NOT EXISTS meta (
		key TEXT PRIMARY KEY,
		value INTEGER
	);
`

const initSchemaVersionSQL = `
	INSERT OR IGNORE INTO meta (key, value) VALUES ('schema_version', 1);
`

func InitDB(db *sql.DB) {
	if err := createSongsTable(db); err != nil {
		log.Fatalf("Failed to create songs table: %v", err)
	}
	if err := createMetaTable(db); err != nil {
		log.Fatalf("Failed to create meta table: %v", err)
	}
	if err := initSchemaVersion(db); err != nil {
		log.Fatalf("Failed to initialize schema version: %v", err)
	}
	// further function calls to create tables
}

func createSongsTable(db *sql.DB) error {
	_, err := db.Exec(createSongsTableSQL)
	return err
}

func createMetaTable(db *sql.DB) error {
	_, err := db.Exec(createMetaTableSQL)
	return err
}

func initSchemaVersion(db *sql.DB) error {
	_, err := db.Exec(initSchemaVersionSQL)
	return err
}

func GetSchemaVersion(db *sql.DB) (int, error) {
	var version int
	err := db.QueryRow(`SELECT value FROM meta WHERE key = 'schema_version'`).Scan(&version)
	return version, err
}
