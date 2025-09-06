package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Gilgalad195/musicdb/db"
	_ "modernc.org/sqlite"
)

func main() {
	dbConn, err := sql.Open("sqlite", "music.db")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	db.InitDB(dbConn)

	log.Println("Database initialized.")

	if err := CheckForMigration(dbConn); err != nil {
		fmt.Printf("error with migration: %v", err)
		return
	}
	fmt.Println("Database version check complete.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Music Database > ")
		hasToken := scanner.Scan()
		if hasToken {
			lowerString := cleanInput(scanner.Text())
			if len(lowerString) == 0 {
				fmt.Println("Enter help if unsure of commands")
			} else {
				cmdInput := lowerString
				if command, exists := commands[cmdInput]; exists {
					err := command.callback(dbConn)
					if err != nil {
						fmt.Println("Error:", err)
					}
				} else {
					fmt.Println("Unknown command")
				}
			}
		}
	}
}

func cleanInput(text string) string {
	cleanText := strings.Fields(strings.ToLower(text))
	return cleanText[0]
}
