package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var id int
var name,domain string

func main() {
	
	
}

// db connection

func connectPostgresDB() *sql.DB{
	connectTo := "user=postgres dbname =databasename password='********' host=localhost port =5432 sslmode = disable"
	db , err := sql.Open("postgres",connectTo)
	if err != nil{
		fmt.Println(err)
	}
	return db
}

