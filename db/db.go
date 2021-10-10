package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Init() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./subs.db")
	if err != nil {
		return nil, err
	}
	//defer db.Close()

	sqlStmt := `
	create table if not exists subs (id integer not null primary key, original text, translate text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	return db, nil
}
