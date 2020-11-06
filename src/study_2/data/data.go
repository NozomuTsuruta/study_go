package data

import (
	"database/sql"
	"log"
)

var Db *sql.DB // グローバル変数

// アプリ起動時に初期化
func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=chatapp sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
