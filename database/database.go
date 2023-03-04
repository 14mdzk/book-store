package database

import (
	"mini_project_1/model"

	"github.com/stretchr/testify/mock"
)

type Database interface {
	All() []model.Book
	FindById(id int) (model.Book, error)
	Store(book model.Book) error
	Remove(id int) error
}

type MockDBClient struct {
	mock.Mock
}

func (m *MockDBClient) All() []model.Book {
	args := m.Called()
	return args.Get(0).([]model.Book)
}

func (m *MockDBClient) FindById(id int) (model.Book, error) {

	return model.Book{}, nil
}

func (m *MockDBClient) Store(book model.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockDBClient) Remove(id int) error {
	args := m.Called(id)

	return args.Error(0)
}
