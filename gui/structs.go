package main

import "database/sql"

type Song struct {
	ID             int
	Title          string
	Composer       sql.NullString
	FirstLine      sql.NullString
	PDFPath        sql.NullString
	LyricSheetPath sql.NullString
	MediaPath      sql.NullString
	ArchiveDate    sql.NullString
}

type Performance struct {
	ID     int
	SongId int
	Date   string
}

type Theme struct {
	ID   int
	Name string
}
