package repositories

import (
	"go-bootcamp/pkg/domain"
	"go-bootcamp/pkg/models"

	"gorm.io/gorm"
)

var database *gorm.DB

type dbBook struct {
	db *gorm.DB
}

func BookDBInstance(d *gorm.DB) domain.IBookRepo {
	database = d
	return &dbBook{
		db: d,
	}
}

func (repo *dbBook) GetBooks(bookID uint) []models.Book {
	var Book []models.Book
	if bookID != 0 {
		database.Where("id = ?", bookID).Find(&Book)
	} else {
		database.Find(&Book)
	}
	return Book
}

func (repo *dbBook) CreateBook(book *models.Book) error {
	if err := database.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (repo *dbBook) UpdateBook(book *models.Book) error {
	if err := database.Save(book).Error; err != nil {
		return err
	}
	return nil
}
func (repo *dbBook) DeleteBook(bookID uint) error {
	var Book models.Book
	if err := database.Where("id = ?", bookID).Delete(&Book).Error; err != nil {
		return err
	}
	return nil
}
