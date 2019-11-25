package database

import (
	"database/sql"
)

// IDatabase is the interface that wraps the basic database method.
//
// GetDatabase get the sql.DB pointer.
// It returns a *sql.DB.
type IDatabase interface {
	GetDatabase() *sql.DB
}
