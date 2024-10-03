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
	handlers.GetContactById(db, 1)

	// newContact := models.Contact{
	// 	Id:    2,
	// 	Name:  "Modificado",
	// 	Email: "modificado@gmail.com",
	// 	Phone: "9992223333",
	// }
	// handlers.UpdateContact(db, newContact)
	// handlers.CreateContact(db, newContact)
	handlers.DeleteContact(db, 4)
	handlers.ListContacts(db)
}
