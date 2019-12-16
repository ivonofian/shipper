package conf

import (
	"database/sql"
	"fmt"

	//sqlite3 connect
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

//ConfigureDB ConfigureDB
func ConfigureDB() {
	var err error
	db, err = sql.Open("sqlite3", "../database/godb.db")
	if err != nil {
		fmt.Printf("Error DB %s", err.Error())
		panic(err.Error())
	}
	db.Exec("create table if not exists product (id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT,sku TEXT,qty INTEGER,created datetime,updated datetime)")

}

//GetDB return db
func GetDB() *sql.DB {
	return db
}
