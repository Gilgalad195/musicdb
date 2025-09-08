package main

import (
	"database/sql"
	"fmt"

	"github.com/Gilgalad195/musicdb/db"
	_ "modernc.org/sqlite"
)

const currentSchemaVersion = 1

func CheckForMigration(dbase *sql.DB) error {
	dbVersion, err := db.GetSchemaVersion(dbase)
	if err != nil {
		return fmt.Errorf("failed to get schema version: %v", err)
	}
	if dbVersion < currentSchemaVersion {
		if err := runMigration(dbase); err != nil {
			return err
		}
		fmt.Printf("database version updated from v%d to v%d:\n", dbVersion, currentSchemaVersion)
		_, err := dbase.Exec(`UPDATE meta SET value = ? WHERE key = 'schema_version'`, currentSchemaVersion)
		if err != nil {
			return fmt.Errorf("failed to update schema version: %v", err)
		}
	}
	return nil
}

func runMigration(dbase *sql.DB) error {
	_, err := dbase.Exec(`ALTER TABLE songs RENAME COLUMN delete_date TO archive_date`)
	return err
}
