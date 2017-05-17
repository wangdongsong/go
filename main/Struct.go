package main

import "fmt"

func main()  {
	var book1 Books
	var book2 Books

	book1.title = "Go lang"
	book1.author = "wds-go"
	book1.subject = "Go learing"
	book1.bookId = 1

	book2.title = "Java lang"
	book2.author = "wds-java"
	book2.subject = "Java Learing"
	book2.bookId = 2

	fmt.Printf("Book1 title : %a\n", book1.title)
}

type Books struct {
	title string
	author string
	subject string
	bookId int
}
