package book

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\n", b.Title)
	fmt.Printf("Author: %s\n", b.Author)
	fmt.Printf("Pages: %d\n", b.Pages)
}

func NewBook(title string, author string, pages int) *Book{
	return &Book{
		Title: title,
		Author: author,
		Pages: pages,
	}
}