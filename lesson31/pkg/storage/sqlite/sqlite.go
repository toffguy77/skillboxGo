package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"skillbox/middleware"
)

func getDB(name string) (*sql.DB, error) {
	file, err := os.Create("database/data/sqlite3.db")
	if err != nil {
		log.Fatal(err)
	}
}

func getDBFile(name string) *os.File {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND, 0666)
	middleware.CheckErr(err)
	return file
}

func CreateTables(db *sql.DB) {
	err := createPersonsTable(db)
	if err != nil {
		log.Fatal(err)
	}
	err = createFriendsTable(db)
	if err != nil {
		log.Fatal(err)
	}
}

func createPersonsTable(db *sql.DB) error {
	persons_table := `CREATE TABLE persons (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "Name" TEXT,
        "Age" INTEGER,
        "Dept" TEXT,;`
	query, err := db.Prepare(persons_table)
	if err != nil {
		return err
	}
	_, err = query.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Table `persons` created successfully!")
	return nil
}

func createFriendsTable(db *sql.DB) error {
	friends_table := `CREATE TABLE friends (
        "Source" INTEGER NOT NULL PRIMARY KEY,
        "Target" INTEGER NOT NULL,;`
	query, err := db.Prepare(friends_table)
	if err != nil {
		return err
	}
	_, err = query.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Table `friends` created successfully!")
	return nil
}
