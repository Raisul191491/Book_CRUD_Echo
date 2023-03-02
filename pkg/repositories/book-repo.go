package repositories

import (
	"go-bootcamp/pkg/domain"
	"go-bootcamp/pkg/models"

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

func (repo *dbBookstore) BookList() []models.Book {
	var AllBooks []models.Book
	database.Find(&AllBooks)
	return AllBooks
}
func (repo *dbBookstore) GetBooks(bookID uint) models.Book {
	var Book models.Book
	database.Where("id = ?", bookID).Find(&Book)
	return Book
}
func (repo *dbBookstore) CreateBook(book *models.Book) error {
	if err := database.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (repo *dbBookstore) UpdateBook(book *models.Book) error {
	if err := database.Save(book).Error; err != nil {
		return err
	}
	return nil
}
func (repo *dbBookstore) DeleteBook(bookID uint) error {
	var Book models.Book
	if err := database.Where("id = ?", bookID).Delete(&Book).Error; err != nil {
		return err
	}
	return nil
}
