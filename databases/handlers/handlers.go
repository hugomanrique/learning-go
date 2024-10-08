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

func GetContactById(db *sql.DB, id int) {
	query := "SELECT * FROM contacts WHERE id = ?"
	row := db.QueryRow(query, id)
	var contact models.Contact
	var valueMail sql.NullString
	err := row.Scan(&contact.Id, &contact.Name, &valueMail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("Id not found")
		}
		log.Fatal(err)
	}
	if valueMail.Valid {
		contact.Email = valueMail.String
	} else {
		contact.Email = "Sin correo"
	}
	log.Println(contact)
}

func CreateContact(db *sql.DB, contact models.Contact) {
	query := "INSERT INTO contacts (name, email, phone) VALUES (?, ?, ?)"
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contact created")
}

func UpdateContact(db *sql.DB, contact models.Contact) {
	query := "UPDATE contacts SET name = ?, email = ?, phone = ? WHERE id = ?"
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contact updated")
}

func DeleteContact(db *sql.DB, id int) {
	query := "DELETE FROM contacts WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contact deleted")
}

func RestarValue() {

}
