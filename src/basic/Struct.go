package basic

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

	printBook(book2)
	fmt.Printf("Book1 title : %s\n", book1.title)
	printPointBook(&book1)
}

/**
 * 指针
 */
func printPointBook(books *Books) {
	fmt.Printf("Book title: %s\n", books.title)
}

/**
 * 函数
 */
func printBook(book Books) {
	fmt.Printf("Book1 title : %s\n", book.title)
}

type Books struct {
	title string
	author string
	subject string
	bookId int
}
