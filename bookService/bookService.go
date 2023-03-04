package bookService

import (
	"mini_project_1/database"
	"mini_project_1/model"
)

type BookService struct {
	DBClient database.Database
}

func (bs *BookService) FetchBook() []model.Book {
	return bs.DBClient.All()
}

func (bs *BookService) AddBook(book model.Book) error {
	err := bs.DBClient.Store(book)

	if err != nil {
		return err
	}

	return nil
}

func (bs *BookService) FindBook(id int) (model.Book, error) {
	book, err := bs.DBClient.FindById(id)
	return book, err
}

func (bs *BookService) RemoveBook(id int) error {
	err := bs.DBClient.Remove(id)
	return err
}
