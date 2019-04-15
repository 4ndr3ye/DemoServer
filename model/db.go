package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type PostGresSQL struct {
	Username string
	Password string
	DBName   string
	Hostname string
	SSLMode  bool
}

const (
	hostname = "localhost"
	username = "testuser"
	password = "Aa123456!"
	dbame    = "dvdrental"
	ssl      = false
)

var (
	SQL       *sql.DB
	Databases PostGresSQL
)

// Connect to the database
func Connect(d PostGresSQL) {
	var err error

	Databases = d

	if SQL, err = sql.Open("postgres", DSN()); err != nil {
		log.Println("SQL Driver Error", err)
	}

	if err = SQL.Ping(); err != nil {
		log.Println("Database Error", err)
	}
}

func DSN() string {
	// Example: postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full
	return "postgres://" +
		username +
		":" +
		password +
		"@" +
		hostname +
		"/" +
		dbame +
		fmt.Sprintf("%s", SSLMode(ssl))
}

func SSLMode(mode bool) string {
	if mode {
		return "?sslmode=verify-full"
	}
	return ""
}
