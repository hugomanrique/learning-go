package book

import "fmt"

type Book struct {
	title  string
	author string
	pages  int
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\n", b.title)
	fmt.Printf("Author: %s\n", b.author)
	fmt.Printf("Pages: %d\n", b.pages)
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetPages() int {
	return b.pages
}

type TextBook struct {
	Book
	editorial string
	level     string
}

func NewTexBook(title string, author string, pages int, editorial string, level string) *TextBook {
	return &TextBook{
		Book: Book{
			title:  title,
			author: author,
			pages:  pages,
		},
		editorial: editorial,
		level:     level,
	}
}

func (tb *TextBook) PrintInfo() {
	fmt.Printf("Title: %s\n", tb.title)
	fmt.Printf("Author: %s\n", tb.author)
	fmt.Printf("Pages: %d\n", tb.pages)
	fmt.Printf("Editorial: %s\n", tb.editorial)
	fmt.Printf("Level: %s\n", tb.level)
}

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}
