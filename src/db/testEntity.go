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

import "database/sql"

type TestEntity struct {
	Text string `json:"text"`
}

func scanTestEntity(rows *sql.Rows) (*TestEntity, error) {
	rows.Next()
	var entity TestEntity
	err := rows.Scan(&entity.Text)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

// Searches for TestEntity with text and returns it
func (db *DB) GetTestEntity(text string) (*TestEntity, error) {
	rows, err := db.Query("SELECT * FROM test WHERE text=?", text)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	TestEntity, err := scanTestEntity(rows)
	if err != nil {
		return nil, err
	}

	return TestEntity, nil
}

// Creates a new TestEntity in the database
func (db *DB) CreateTestEntity(newTestEntity TestEntity) error {
	_, err := db.Exec(
		"INSERT INTO test(text) VALUES(?)",
		newTestEntity.Text,
	)

	return err
}

// Deletes TestEntity with given text
func (db *DB) DeleteTestEntity(text string) error {
	_, err := db.Exec(
		"DELETE FROM test WHERE text=?",
		text,
	)

	return err
}
