package pkg

import "strconv"

var bookId int = 0
var db = map[string]Book{
	strconv.Itoa(bookId): {
		Name:        "The Alchemist",
		Description: "xyz",
		Author:      "pqr",
		Price:       200,
	}}

func ReadBooksFromDB() *[]Book {
	var books []Book
	for _, book := range db {
		books = append(books, book)
	}
	return &books
}

func CreateBookInDB(book *Book) {
	bookId += 1
	id := strconv.Itoa(bookId)
	db[id] = *book
}

func ReadBookByIdFromDB(id string) *Book {
	if book, iskey := db[id]; iskey {
		return &book
	}
	return nil
}
