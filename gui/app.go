package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Gilgalad195/musicdb/db"
)

// App struct
type App struct {
	DBase *sql.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	dbConn, err := sql.Open("sqlite", "music.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.InitDB(dbConn); err != nil {
		log.Fatalf("DB init failed: %v", err)
	}

	if err := CheckForMigration(dbConn); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database initialized and migrated.")

	return &App{DBase: dbConn}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
