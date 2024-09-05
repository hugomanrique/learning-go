package main

import "library/book"

func main() {
	// var book book.Book = book.Book{
	// 	Title:  "The Hobbit",
	// 	Author: "J.R.R. Tolkien",
	// 	Pages:  1024,
	// }
	// book.PrintInfo()

	myBook := book.NewBook("The Hobbit", "J.R.R. Tolkien", 1024)
	myBook.PrintInfo()
}