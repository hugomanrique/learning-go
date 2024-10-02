package main

import (
	"log"
	"mysql-go/db"
	"mysql-go/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	handlers.ListContacts(db)
}
