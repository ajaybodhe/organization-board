package db

import (
	"database/sql"
	"log"

	"personio.com/organization-board/config"

	_ "github.com/mattn/go-sqlite3"
)

const (
	sqliteDriverName = "sqlite3"
)

var queries = []string{
	`CREATE TABLE IF NOT EXISTS user_detail (
	    id                  INTEGER     PRIMARY KEY,
	    email               TEXT        NOT NULL UNIQUE,
	    password            TEXT        NOT NULL,
	    deleted             INT         NOT NULL DEFAULT 0
	);`,

	`INSERT OR IGNORE INTO user_detail (email, password, deleted) VALUES ('personia@org.com', 'personia', 0)`,
}

type SQLLite struct {
}

func (lite *SQLLite) NewConnection() (*sql.DB, error) {
	sqliteCfg := config.Config().SQLite
	log.Println(sqliteCfg.DataSourceName)
	return sql.Open(sqliteDriverName, sqliteCfg.DataSourceName)
}

func init() {
	sqlite := new(SQLLite)
	conn, err := sqlite.NewConnection()
	if err != err {
		log.Fatalf("Error while creating SQLLite connection:%s", err.Error())
	}

	for _, query := range queries {
		if _, err := conn.Exec(query); nil != err {
			log.Fatalf("Error while executing query:%s:%s", query, err.Error())
		}
	}

	if err := conn.Close(); nil != err {
		log.Fatalf("Error while closing SQLLite connection:%s", err.Error())
	}

	log.Println("SQLite DB initialised successfully")
}
