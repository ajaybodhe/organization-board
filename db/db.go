package db

import (
	"database/sql"
)

// IDB : Interface to perform DB operations.
// TODO : support dependancy injection
type IDB interface {
	NewConnection() (*sql.DB, error)
}
