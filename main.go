package main

import (
	"fmt"
	"mini_project_1/bookService"
	"mini_project_1/database"
	"mini_project_1/model"
)

func main() {
	bookService := bookService.BookService{
		DBClient: &database.MemoryDBClient{},
	}

	fmt.Println("Fetch book data: \n", bookService.FetchBook())

	bookService.AddBook(model.Book{
		BookId:        1001,
		Title:         "GO 101",
		Author:        "Unknown",
		PublishedYear: 2020,
		Quantity:      5,
		IsAvailable:   true,
	})

	bookService.AddBook(model.Book{
		BookId:        1002,
		Title:         "UNDERSTANDING GO 502",
		Author:        "Unknown",
		PublishedYear: 2021,
		Quantity:      2,
		IsAvailable:   true,
	})

	fmt.Println("Fetch book data: \n", bookService.FetchBook())

	removeBook := bookService.RemoveBook(1004)
	if removeBook == nil {
		fmt.Println("Successfully removing book with ID ", 1004)
	} else {
		fmt.Println("Failed to remove book.", removeBook)
	}

	removeBook = bookService.RemoveBook(1002)
	if removeBook == nil {
		fmt.Println("Successfully removing book with ID ", 1002)
	} else {
		fmt.Println("Failed to remove book.", removeBook)
	}

	fmt.Println("Fetch book data: \n", bookService.FetchBook())

}
