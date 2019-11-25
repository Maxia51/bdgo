package database

import (
	"database/sql"
)

type IDatabase interface {
	GetDatabase() *sql.DB
}
