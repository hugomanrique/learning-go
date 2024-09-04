package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("holal mundo")
	fmt.Println(quote.Hello())

	var name, name2 string
	var age int
	fmt.Println("Digita tu nombre")
	fmt.Scanln(&name,&name2)
	fmt.Println("Digita tu edad")
	fmt.Scanln(&age)
	fmt.Printf("Hola me llamo %s %s y tengo %d años \n", name,name2, age)

	greeting := fmt.Sprintf("Hola me llamo %v %v y tengo %v años \n", name,name2, age)
	fmt.Println(greeting)
}
