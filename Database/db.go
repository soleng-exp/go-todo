package Database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Db *sql.DB
)

func InitDB(filepath string) {
	Db = initDB(filepath)
	migrate(Db)
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
    );
    `

	_, err := db.Exec(query)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
