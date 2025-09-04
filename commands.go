package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func prompt(scanner *bufio.Scanner, systemPrompt string) string {
	fmt.Printf("%s ", systemPrompt)
	var response string
	hasToken := scanner.Scan()
	if hasToken {
		response = strings.TrimSpace(scanner.Text())
	}
	return response
}

func nullToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return "(none)"
}

func commandExit(_ *sql.DB) error {
	fmt.Println("Closing the Database... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *sql.DB) error {
	fmt.Println("Welcome to the Music Libary!")
	fmt.Printf("Usage:\n\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandAdd(db *sql.DB) error {
	fmt.Println("Adding song to library")
	scanner := bufio.NewScanner(os.Stdin)
	title := prompt(scanner, "Enter the song title: ")
	if title == "" {
		fmt.Println("A title is required")
		return nil
	}
	composer := prompt(scanner, "Enter the composer(s) or song writer(s): ")
	_, err := db.Exec("INSERT INTO songs (title, composer) VALUES (?, ?)", title, composer)
	if err != nil {
		return fmt.Errorf("error adding song: %v", err)
	}
	fmt.Printf("%s by %s added to music library\n", title, composer)
	return nil
}

func commandList(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM songs")
	if err != nil {
		return fmt.Errorf("error retrieving songs: %v", err)
	}
	fmt.Println("")
	for rows.Next() {
		var (
			id               int
			title            string
			composer         sql.NullString
			first_line       sql.NullString
			themes           sql.NullString
			scripture_refs   sql.NullString
			pdf_path         sql.NullString
			lyric_sheet_path sql.NullString
			media_path       sql.NullString
		)
		err := rows.Scan(&id, &title, &composer, &first_line, &themes, &scripture_refs, &pdf_path, &lyric_sheet_path, &media_path)
		if err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}
		fmt.Printf("ID: %d\nTitle: %s\nComposer: %s\nFirst Line: %s\nThemes: %s\nScripture Refs: %s\nPDF: %s\nLyric Sheet: %s\nMedia: %s\n\n",
			id, title, nullToString(composer), nullToString(first_line), nullToString(themes), nullToString(scripture_refs), nullToString(pdf_path), nullToString(lyric_sheet_path), nullToString(media_path))
	}
	return nil
}
