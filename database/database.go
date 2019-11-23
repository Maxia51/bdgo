package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type service struct {
	db *sql.DB
}

// New instanciate a database struct
// it return a database struct
func New() IDatabase {

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_USER_PASSWORD")+"@tcp("+os.Getenv("MYSQL_ADDR")+":"+os.Getenv("MYSQL_PORT")+")/"+os.Getenv("MYSQL_DATABASE_NAME")+"?parseTime=true")

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
