package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

func commandExit(_ *sql.DB, _ []string) error {
	fmt.Println("Closing the Database... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *sql.DB, _ []string) error {
	fmt.Println("Welcome to the Music Libary!")
	fmt.Printf("Usage:\n\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandAdd(db *sql.DB, _ []string) error {
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

// need to update this to list performance data with --log
func commandList(dbase *sql.DB, args []string) error {
	includeArchived := len(args) > 0 && args[0] == "--admin"
	query := `SELECT id, title, composer, first_line, pdf_path, lyric_sheet_path, media_path, archive_date FROM songs`
	if !includeArchived {
		query += " WHERE archive_date IS NULL"
	}
	rows, err := dbase.Query(query)
	if err != nil {
		return fmt.Errorf("error retrieving songs: %v", err)
	}
	defer rows.Close()
	fmt.Println("")
	for rows.Next() {
		var (
			id               int
			title            string
			composer         sql.NullString
			first_line       sql.NullString
			pdf_path         sql.NullString
			lyric_sheet_path sql.NullString
			media_path       sql.NullString
			archive_date     sql.NullString
		)
		err := rows.Scan(&id, &title, &composer, &first_line, &pdf_path, &lyric_sheet_path, &media_path, &archive_date)
		if err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}
		fmt.Printf("ID: %d\nTitle: %s\nComposer: %s\nFirst Line: %s\nPDF: %s\nLyric Sheet: %s\nMedia: %s\n",
			id, title, nullToString(composer), nullToString(first_line), nullToString(pdf_path), nullToString(lyric_sheet_path), nullToString(media_path))
		if includeArchived {
			fmt.Printf("Archive Date: %s\n", nullToString(archive_date))
		}
		fmt.Println("")
	}
	return nil
}

func commandArchive(dbase *sql.DB, _ []string) error {
	scanner := bufio.NewScanner(os.Stdin)
	songIdString := prompt(scanner, "Enter the Song ID to archive:")
	songId, err := strconv.Atoi(songIdString)
	if err != nil {
		return fmt.Errorf("song ID must be a numeric value: %v", err)
	}
	date := time.Now().Format("2006-01-02")
	if _, err := dbase.Exec(`UPDATE songs SET archive_date = ? WHERE id = ?`, date, songId); err != nil {
		return fmt.Errorf("failed to archive song: %v", err)
	}
	fmt.Println("Song successfully archived")
	return nil
}

func commandUnarchive(dbase *sql.DB, _ []string) error {
	scanner := bufio.NewScanner(os.Stdin)
	songIdString := prompt(scanner, "Enter the Song ID to unarchive:")
	songId, err := strconv.Atoi(songIdString)
	if err != nil {
		return fmt.Errorf("song ID must be a numeric value: %v", err)
	}
	if _, err := dbase.Exec(`UPDATE songs SET archive_date = NULL WHERE id = ?`, songId); err != nil {
		return fmt.Errorf("failed to unarchive song: %v", err)
	}
	fmt.Println("Song successfully unarchived")
	return nil
}

func commandLog(dbase *sql.DB, args []string) error {
	scanner := bufio.NewScanner(os.Stdin)
	if len(args) > 0 && args[0] == "--del" {
		perfIdString := prompt(scanner, "Enter the Performance ID to delete:")
		perfId, err := strconv.Atoi(perfIdString)
		if err != nil {
			return fmt.Errorf("not a valid performance id: %v", err)
		}
		confirm := prompt(scanner, fmt.Sprintf("Are you sure you want to delete performance #%d? (y/n):", perfId))
		if strings.ToLower(confirm) != "y" {
			fmt.Println("Deletion cancelled.")
			return nil
		}
		if _, err := dbase.Exec("DELETE FROM performances WHERE id = ?", perfId); err != nil {
			return fmt.Errorf("error deleting performance: %v", err)
		}
		fmt.Println("Performance record deleted.")
		return nil
	}
	songIdString := prompt(scanner, "Enter the Song ID to log:")
	songId, err := strconv.Atoi(songIdString)
	if err != nil {
		return fmt.Errorf("not a valid song id: %v", err)
	}
	date := prompt(scanner, "Enter the performance date (YYYY-MM-DD):")
	if _, err := dbase.Exec("INSERT INTO performances (song_id, date) VALUES (?, ?)", songId, date); err != nil {
		return fmt.Errorf("error adding song: %v", err)
	}
	fmt.Printf("Song Id %d logged for %s\n", songId, date)
	return nil
}
