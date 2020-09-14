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

// Dummy data for authenticated user
var queries = []string{
	`CREATE TABLE IF NOT EXISTS user_detail (
	    id                  INTEGER     PRIMARY KEY,
	    email               TEXT        NOT NULL UNIQUE,
	    password            TEXT        NOT NULL,
	    deleted             INT         NOT NULL DEFAULT 0
	);`,

	// TODO : we can create seeder package for inserting data in initial state
	`INSERT OR IGNORE INTO user_detail (email, password, deleted) VALUES ('personia@org.com', 'personia', 0);`,

	`CREATE TABLE IF NOT EXISTS employee_manager_mapping (
	    employee_name       TEXT        NOT NULL PRIMARY KEY,
	    manager_name        TEXT        NOT NULL
	);`,
}

// SQLLite : implement IDB interface for sqllite DB
type SQLLite struct {
}

// NewConnection : new sqllite conneection struct
func (lite *SQLLite) NewConnection() (*sql.DB, error) {
	sqliteCfg := config.Config().SQLite
	log.Println(sqliteCfg.DataSourceName)
	return sql.Open(sqliteDriverName, sqliteCfg.DataSourceName)
}

// we are facillitating user authentication through dummy user
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
