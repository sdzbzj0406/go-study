package main

import (
	"fmt"
)

type myInt int

type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	book.auth = "auth2"
}

func changeBook2(book *Book) {
	book.auth = "auth2"
}

func main2() {

	var a myInt = 10
	fmt.Println(a)

	var book1 Book
	book1.title = "title"
	book1.auth = "auth"
	fmt.Println(book1)

	changeBook(book1)
	fmt.Println(book1)

	changeBook2(&book1)
	fmt.Println(book1)
}
