package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type service struct {
	db *sql.DB
}

// New instanciate a database struct
// it return a database struct
func New(user string, password string, ip string, port string, name string) IDatabase {

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+ip+":"+port+")/"+name)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return &service{
		db: db,
	}
}

func (s *service) GetDatabase() *sql.DB {
	return s.db
}
