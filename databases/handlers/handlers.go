package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"mysql-go/models"
)

func ListContacts(db *sql.DB) {
	query := "SELECT * FROM contacts"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("List of contacts:")
	var valueMail sql.NullString
	for rows.Next() {
		var contact models.Contact
		err := rows.Scan(&contact.Id, &contact.Name, &valueMail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		if valueMail.Valid {
			contact.Email = valueMail.String
		} else {
			contact.Email = "Sin correo"
		}
		fmt.Println(contact)
	}
}
