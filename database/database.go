package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "8426" // TODO: change to env variable
	dbname   = "crowd-compute"
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}

	fmt.Println("<-- Database connection established -->")
}
