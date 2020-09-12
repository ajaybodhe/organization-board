package db

import (
	"database/sql"

	"personio.com/organization-board/config"
)

const (
	sqliteDriverName = "sqlite3"
)

type SQLLite struct {
}

func (lite *SQLLite) NewConnection() (*sql.DB, error) {
	sqliteCfg := config.Config().SQLite
	return sql.Open(sqliteDriverName, sqliteCfg.DataSourceName)
}
