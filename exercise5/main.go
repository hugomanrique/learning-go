package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"hugomanrique.com/greetings"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return a / b, nil
}

func main() {
	// str := "123f"
	// nm, error := strconv.Atoi(str)
	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }
	// fmt.Println(nm)
	// result, error := divide(10, 2)
	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }
	// fmt.Println(result)

	// panicFunction()
	// registerError()
	// contacts()
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // Quit datetime of log

	message, error := greetings.Hello("Andres")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(message)

	names := []string{"Andres", "Hugo", "Maria", "Pedro"}
	messages, error := greetings.HellosMultiple(names)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(messages)
}

// defer
func fn1() {
	file, error := os.Create("test.txt")
	if error != nil {
		fmt.Println(error)
		return
	}

	defer file.Close() // pospouse execution function befoer ends of fn1

	_, error = file.WriteString("hello world")
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println("fn1")
	defer fmt.Println("defer")
}

// panic
func panicFunction() {
	result, _ := divide2(10, 2)
	fmt.Println(result)
	result2, _ := divide2(10, 5)
	fmt.Println(result2)
	resul4, _ := divide2(10, 0)
	fmt.Println(resul4)
	resul3, _ := divide2(10, 1)
	fmt.Println(resul3)

}

func divide2(a, b int) (int, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered", r)
		}
	}()
	validateZero(b)
	return a / b, nil
}

func validateZero(a int) {
	if a == 0 {
		panic("Can't divide by zero")
	}
}

func registerError() {
	log.Fatal("Fatal")
	log.Println("Message register")
}

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func saveContactsToFile(contacts []Contact) error {
	file, error := os.Create("contacts.json")
	if error != nil {
		return error
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	error = encoder.Encode(contacts)
	if error != nil {
		return error
	}
	return nil
}

func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}

func contacts() {
	var contacts []Contact
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("Choose action:")
		fmt.Println("1. Add contact")
		fmt.Println("2. Show contacts")
		fmt.Println("3. Exit")

		fmt.Print("Enter action: ")
		var action int
		_, err := fmt.Scanln(&action)
		if err != nil {
			fmt.Println("Invalid action")
			return
		}

		switch action {
		case 1:
			var ct Contact
			fmt.Print("Enter name: ")
			ct.Name, _ = reader.ReadString('\n')
			fmt.Print("Enter phone: ")
			ct.Phone, _ = reader.ReadString('\n')
			fmt.Print("Enter email: ")
			ct.Email, _ = reader.ReadString('\n')
			contacts = append(contacts, ct)
			fmt.Println("Contact added")

			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println(err)
				return
			}
		case 2:
			fmt.Println("Contacts:")
			for _, contact := range contacts {
				fmt.Println(contact.Name, contact.Phone, contact.Email)
			}
			fmt.Println("+++++++++++++++++++++++++++++++++++")
		case 3:
			fmt.Println("Bye")
			return
		default:
			fmt.Println("Invalid action")
		}

	}
}
