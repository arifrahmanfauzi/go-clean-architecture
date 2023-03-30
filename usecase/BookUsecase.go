package usecase

import (
	"go-clean-architecture/models"
	"go-clean-architecture/repository"
)

type BookUsecaseInterface interface {
	Fetch() []models.Book
	GetById(id int64) (*models.Book, error)
	GetByTitle(title string) (*models.Book, error)
	Update(book *models.Book) (models.Book, error)
	Store(book *models.BookRequest) (models.Book, error)
	Delete(id int64) (bool, error)
}

type BookUsecase struct {
	BookRepository repository.BookRepositoryContract
}

func NewBookUsecase(r repository.BookRepositoryContract) *BookUsecase {
	return &BookUsecase{BookRepository: r}
}
func (bu *BookUsecase) Fetch() []models.Book {
	res := bu.BookRepository.Fetch()
	return res
}

func (bu *BookUsecase) GetById(id int64) (*models.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (bu *BookUsecase) GetByTitle(title string) (*models.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (bu *BookUsecase) Update(input *models.Book) (models.Book, error) {

	res, err := bu.BookRepository.Update(input)
	return res, err
}

func (bu *BookUsecase) Delete(id int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (bu *BookUsecase) Store(input *models.BookRequest) (models.Book, error) {
	var books models.Book
	books.Title = input.Title
	books.Author = input.Author
	res, err := bu.BookRepository.Store(&books)
	return res, err
}
