package models

import (
	"database/sql"
	"fmt"
	"go_crud/config"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	DB, err = sql.Open(config.Config.SQLDriver, config.Config.DBName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING NOT NULL,
		email STRING NOT NULL UNIQUE,
		password STRING NOT NULL,
		created_at DATETIME NOT NULL)`, tableNameUser)

	DB.Exec(cmdU)
}
