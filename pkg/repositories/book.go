package repositories

import (
	"go-bootcamp/pkg/domain"
	"go-bootcamp/pkg/models"

	"gorm.io/gorm"
)

type dbBook struct {
	db *gorm.DB
}

func BookDBInstance(d *gorm.DB) domain.IBookRepo {
	return &dbBook{
		db: d,
	}
}

func (dbObject *dbBook) GetBooks(bookID uint) []models.Book {
	var Book []models.Book
	if bookID != 0 {
		dbObject.db.Where("id = ?", bookID).Find(&Book)
	} else {
		dbObject.db.Find(&Book)
	}
	return Book
}

func (dbObject *dbBook) CreateBook(book *models.Book) error {
	if err := dbObject.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (dbObject *dbBook) UpdateBook(book *models.Book) error {
	if err := dbObject.db.Save(book).Error; err != nil {
		return err
	}
	return nil
}
func (dbObject *dbBook) DeleteBook(bookID uint) error {
	var Book models.Book
	if err := dbObject.db.Where("id = ?", bookID).Delete(&Book).Error; err != nil {
		return err
	}
	return nil
}
