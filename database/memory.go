package database

import (
	"errors"
	"fmt"
	"mini_project_1/model"
)

type MemoryDBClient struct {
	Books []model.Book
}

func (mr *MemoryDBClient) All() []model.Book {
	return mr.Books
}

func (mr *MemoryDBClient) Store(book model.Book) error {
	if (book == model.Book{}) {
		return errors.New("Book shouldn't be empty")
	}

	mr.Books = append(mr.Books, book)

	return nil
}

func (mr *MemoryDBClient) Remove(id int) error {
	_, err := mr.FindById(id)

	for key, item := range mr.Books {
		if item.BookId == id {
			mr.Books = append(mr.Books[:key], mr.Books[key+1:]...)
			break
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func (mr *MemoryDBClient) FindById(id int) (model.Book, error) {
	var book model.Book

	for _, item := range mr.Books {
		if item.BookId == id {
			book = item
			break
		}
	}

	if (book == model.Book{}) {
		err := errors.New(fmt.Sprint("Book with ID ", id, " was not found."))
		return book, err
	}

	return book, nil
}
