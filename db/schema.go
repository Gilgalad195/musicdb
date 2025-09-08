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
		pdf_path TEXT,
		lyric_sheet_path TEXT,
		media_path TEXT,
		archive_date TEXT
	);
`

const createPerformancesTableSQL = `
	CREATE TABLE IF NOT EXISTS performances (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		song_id INTEGER NOT NULL,
		date TEXT NOT NULL,
		FOREIGN KEY (song_id) REFERENCES songs(id)
	);
`

const createThemesTableSQL = `
	CREATE TABLE IF NOT EXISTS themes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);
`

const createSongThemeTableSQL = `
	CREATE TABLE IF NOT EXISTS song_theme (
		song_id INTEGER NOT NULL,
		theme_id INTEGER NOT NULL,
		FOREIGN KEY (song_id) REFERENCES songs(id),
		FOREIGN KEY (theme_id) REFERENCES themes(id),
		PRIMARY KEY (song_id, theme_id)
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

func InitDB(dbase *sql.DB) {
	if err := createSongsTable(dbase); err != nil {
		log.Fatalf("Failed to create songs table: %v", err)
	}
	if err := createPerformancesTable(dbase); err != nil {
		log.Fatalf("Failed to create performances table: %v", err)
	}
	if err := createThemesTable(dbase); err != nil {
		log.Fatalf("Failed to create themes table: %v", err)
	}
	if err := createSongThemeTable(dbase); err != nil {
		log.Fatalf("Failed to create song_theme table: %v", err)
	}
	if err := createMetaTable(dbase); err != nil {
		log.Fatalf("Failed to create meta table: %v", err)
	}
	if err := initSchemaVersion(dbase); err != nil {
		log.Fatalf("Failed to initialize schema version: %v", err)
	}
	// further function calls to create tables
}

func createSongsTable(dbase *sql.DB) error {
	_, err := dbase.Exec(createSongsTableSQL)
	return err
}

func createPerformancesTable(dbase *sql.DB) error {
	_, err := dbase.Exec(createPerformancesTableSQL)
	return err
}

func createThemesTable(dbase *sql.DB) error {
	_, err := dbase.Exec(createThemesTableSQL)
	return err
}

func createSongThemeTable(dbase *sql.DB) error {
	_, err := dbase.Exec(createSongThemeTableSQL)
	return err
}

func createMetaTable(dbase *sql.DB) error {
	_, err := dbase.Exec(createMetaTableSQL)
	return err
}

func initSchemaVersion(dbase *sql.DB) error {
	_, err := dbase.Exec(initSchemaVersionSQL)
	return err
}

func GetSchemaVersion(dbase *sql.DB) (int, error) {
	var version int
	err := dbase.QueryRow(`SELECT value FROM meta WHERE key = 'schema_version'`).Scan(&version)
	return version, err
}
