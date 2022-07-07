package models

import (
	"errors"
	"fmt"

	"github.com/geo-albin/bookStore/pkg/db"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Author string `json:"author"`
}

func init() {
	err := db.Connect()
	if err != nil {
		panic(err)
	}

	db.DB().AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.DB().Create(&b)
	return b
}

func GetAllBooks() ([]Book, error) {
	var books []Book
	err := db.DB().Find(&books).Error
	return books, err
}

func GetBookById(id int64) (Book, error) {
	var book Book
	result := db.DB().Where("ID=?", id).Find(&book)
	err := result.Error

	if result.RowsAffected == 0 {
		l := fmt.Sprintf("Didnt get the book by id %v", id)
		err = errors.New(l)
	}

	return book, err
}

func DeleteBook(id int64) error {
	var book Book
	err := db.DB().Where("ID=?", id).Delete(&book).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}

func GetNextId() (int64, error) {
	var book Book
	var id int64 = 1
	err := db.DB().Last(&book).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return id, err
	}

	id = (int64)(book.ID + 1)

	return id, nil
}
