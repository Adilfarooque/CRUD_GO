package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
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

func insert(db *sql.DB){
	fmt.Println("Enter the id:")
	fmt.Scan(&id)
	fmt.Println("Enter the name:")
	fmt.Scan(&name)
	fmt.Println("Enter the domain:")
	fmt.Scan(&domain)
	insertIntoPostgres(db,id,name,domain)
}

func insertIntoPostgres(db *sql.DB,id int,name string,domain string){
	_,err :=db.Exec("INSERT INTO student(id,name,domain)value($1,$2,$3)",id,name,domain)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Value Inserted!!")
	}
}

func Read(db *sql.DB){
	rows,err := db.Query("SELECT * FROM student")
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("id name domain")
		for rows.Next() {
			rows.Scan(&id,&name,&domain)
			fmt.Println("%d - %s - %s \n",id,name,domain)
		}
	}
}


