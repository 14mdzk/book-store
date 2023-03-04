package bookService

import (
	"errors"
	"mini_project_1/database"
	"mini_project_1/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchBook_normalFlow(t *testing.T) {
	var expectedBooks []model.Book

	dbClientMock := database.MockDBClient{}
	dbClientMock.On("All").Return(expectedBooks)

	bookService := BookService{DBClient: &dbClientMock}
	books := bookService.FetchBook()

	assert.Equal(t, expectedBooks, books)
}

func TestAddBook_normalFlow(t *testing.T) {
	expectedBook := model.Book{
		BookId:        1001,
		Title:         "GO 101",
		Author:        "Unknown",
		PublishedYear: 2020,
		Quantity:      5,
		IsAvailable:   true,
	}

	dbClientMock := database.MockDBClient{}
	dbClientMock.On("Store", expectedBook).Return(nil)

	bookService := BookService{DBClient: &dbClientMock}
	err := bookService.AddBook(expectedBook)

	assert.Equal(t, err, nil)
}

func TestAddBook_bookEmpty(t *testing.T) {
	expectedBook := model.Book{}

	dbClientMock := database.MockDBClient{}
	dbClientMock.On("Store", mock.Anything).Return(errors.New("Book shouldn't be empty"))

	bookService := BookService{DBClient: &dbClientMock}
	err := bookService.AddBook(expectedBook)

	assert.Error(t, err)
}
