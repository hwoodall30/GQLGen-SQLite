package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DataBase struct {
	DBConnection *sqlx.DB
}

func (d *DataBase) InitDatabase() {
	var err error
	d.DBConnection, err = sqlx.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatal(err)
	}
	err = d.DBConnection.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the SQLite database")
}
