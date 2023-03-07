package repositories

import (
	"go-bootcamp/pkg/app/domain"

	"go-bootcamp/pkg/infra/connection/db/model"

	"gorm.io/gorm"
)

var database *gorm.DB

type dbBookstore struct {
	db *gorm.DB
}

func BookStoreDBInstance(d *gorm.DB) domain.IBookstoreRepo {
	database = d
	return &dbBookstore{
		db: d,
	}
}

func (repo *dbBookstore) GetBooks(bookID uint) []model.Book {
	var Book []model.Book
	if bookID != 0 {
		database.Where("id = ?", bookID).Find(&Book)
	} else {
		database.Find(&Book)
	}
	return Book
}

func (repo *dbBookstore) CreateBook(book *model.Book) error {
	if err := database.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (repo *dbBookstore) UpdateBook(book *model.Book) error {
	if err := database.Save(book).Error; err != nil {
		return err
	}
	return nil
}
func (repo *dbBookstore) DeleteBook(bookID uint) error {
	var Book model.Book
	if err := database.Where("id = ?", bookID).Delete(&Book).Error; err != nil {
		return err
	}
	return nil
}
