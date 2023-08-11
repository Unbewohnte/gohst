/*
            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                    Version 2, December 2004

 Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>

 Everyone is permitted to copy and distribute verbatim or modified
 copies of this license document, and changing it is allowed as long
 as the name is changed.

            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

  0. You just DO WHAT THE FUCK YOU WANT TO.
*/

// Kasyanov N.A. (Unbewohnte), 2023

package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite" // For example this one
)

// Database wrapper
type DB struct {
	*sql.DB
}

func setUpTables(db *DB) error {
	// Table for test entities to be stored in
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS test(id INTEGER PRIMARY KEY, data TEXT NOT NULL)`)
	if err != nil {
		return err
	}

	return nil
}

// Open database
func FromFile(path string) (*DB, error) {
	driver, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	dbase := &DB{driver}

	err = setUpTables(dbase)
	if err != nil {
		return nil, err
	}

	return dbase, nil
}

// Create database file
func Create(path string) (*DB, error) {
	dbFile, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	dbFile.Close()

	driver, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	dbase := &DB{driver}

	err = setUpTables(dbase)
	if err != nil {
		return nil, err
	}

	return dbase, nil
}
