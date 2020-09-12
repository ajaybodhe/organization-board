package db

import (
	"database/sql"
)

type IDB interface {
	NewConnection() (*sql.DB, error)
}
