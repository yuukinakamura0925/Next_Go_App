package usecases

import (
	"Next_Go_App/ent"
	"Next_Go_App/internal/repositories"
	"context"
)

type IBookUsecase interface {
	CreateBook(ctx context.Context, title string, body string, userID int) (*ent.Book, error)
	GetAllBooks(ctx context.Context) ([]*ent.Book, error)
	GetBookByID(ctx context.Context, id int) (*ent.Book, error)
	UpdateBook(ctx context.Context, id int, title string, body string, userID int) (*ent.Book, error)
	DeleteBook(ctx context.Context, id int) error
}

type bookUsecase struct {
	bookRepo repositories.IBookRepository
}

func NewBookUsecase(bookRepo repositories.IBookRepository) IBookUsecase {
	return &bookUsecase{bookRepo: bookRepo}
}

func (u *bookUsecase) CreateBook(ctx context.Context, title string, body string, userID int) (*ent.Book, error) {
	return u.bookRepo.CreateBook(ctx, title, body, userID)
}

func (u *bookUsecase) GetAllBooks(ctx context.Context) ([]*ent.Book, error) {
	return u.bookRepo.GetAllBooks(ctx)
}

func (u *bookUsecase) GetBookByID(ctx context.Context, id int) (*ent.Book, error) {
	return u.bookRepo.GetBookByID(ctx, id)
}

func (u *bookUsecase) UpdateBook(ctx context.Context, id int, title string, body string, userID int) (*ent.Book, error) {
	return u.bookRepo.UpdateBook(ctx, id, title, body, userID)
}

func (u *bookUsecase) DeleteBook(ctx context.Context, id int) error {
	return u.bookRepo.DeleteBook(ctx, id)
}
