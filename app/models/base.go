package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"go_crud/config"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

var err error

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
	tableNameSession = "sessions"
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

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	DB.Exec(cmdT)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING NOT NULL,
		user_id INTEGER NOT NULL,
		created_at DATETIME NOT NULL)`, tableNameSession)

	DB.Exec(cmdS)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
