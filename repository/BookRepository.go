package repository

import (
	"go-clean-architecture/models"
	"gorm.io/gorm"
)

type BookRepositoryContract interface {
	Fetch() []models.Book
	Store(book *models.Book) (models.Book, error)
	Update(input *models.Book) (models.Book, error)
}

type BookRepository struct {
	DB *gorm.DB
}

func ProviderBookRepository(DB *gorm.DB) *BookRepository {
	return &BookRepository{DB: DB}
}
func (m *BookRepository) Fetch() []models.Book {
	var books []models.Book
	db := m.DB.Model(&books)
	db.Debug().Find(&books)
	return books
}

func (m BookRepository) Store(book *models.Book) (models.Book, error) {
	db := m.DB.Model(&book)

	if err := db.Debug().Create(&book).Error; err != nil {
		return *book, err
	}
	return *book, nil
}

func (m BookRepository) Update(input *models.Book) (models.Book, error) {
	var books models.Book
	db := m.DB.Model(&books)
	books.ID = input.ID

	checkStudentId := db.Debug().First(&books)
	if checkStudentId.RowsAffected < 1 {
		return books, checkStudentId.Error
	}
	db.Updates(&input)

	return books, nil
}
