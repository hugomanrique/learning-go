package main

import (
	"library/animal"
)

func main() {
	// var book book.Book = book.Book{
	// 	Title:  "The Hobbit",
	// 	Author: "J.R.R. Tolkien",
	// 	Pages:  1024,
	// }
	// book.PrintInfo()

	// myBook := book.NewBook("The Hobbit", "J.R.R. Tolkien", 1024)
	// // myBook.PrintInfo()
	// myBook.SetTitle("The Lord of the Rings (Ultimate edition)")
	// myBook.SetAuthor("J.R.R. Tolkien")
	// // myBook.PrintInfo()

	// texBook := book.NewTexBook("The Hobbit", "J.R.R. Tolkien", 1024, "The Hobbit", "1")
	// // texBook.PrintInfo()

	// book.Print(myBook)
	// book.Print(texBook)

	// dog := animal.Dog{Name: "M&M"}
	// cat := animal.Cat{Name: "Mimi"}
	// animal.DoSound(&dog)
	// animal.DoSound(&cat)

	animales := []animal.Animal{
		&animal.Dog{Name: "M&M"},
		&animal.Cat{Name: "Mimi"},
		&animal.Dog{Name: "Rambo"},
		&animal.Cat{Name: "michi"},
	}

	for _, anim := range animales {
		animal.DoSound(anim)
	}

}
