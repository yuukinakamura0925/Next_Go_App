package repositories

import (
	"Next_Go_App/ent"
	"Next_Go_App/ent/book"
	"context"
)

type IBookRepository interface {
	CreateBook(ctx context.Context, title string, body string, userID int) (*ent.Book, error)
	GetAllBooks(ctx context.Context) ([]*ent.Book, error)
	GetBookByID(ctx context.Context, id int) (*ent.Book, error)
	UpdateBook(ctx context.Context, id int, title string, body string, userID int) (*ent.Book, error)
	DeleteBook(ctx context.Context, id int) error
}

type bookRepository struct {
	client *ent.Client
}

func NewBookRepository(client *ent.Client) IBookRepository {
	return &bookRepository{client: client}
}

func (r *bookRepository) CreateBook(ctx context.Context, title string, body string, userID int) (*ent.Book, error) {
	newBook, err := r.client.Book.
		Create().
		SetTitle(title).
		SetBody(body).
		SetUserID(userID).
		Save(ctx)
	return newBook, err
}

func (r *bookRepository) GetAllBooks(ctx context.Context) ([]*ent.Book, error) {
	books, err := r.client.Book.Query().All(ctx)
	return books, err
}

func (r *bookRepository) GetBookByID(ctx context.Context, id int) (*ent.Book, error) {
	book, err := r.client.Book.Query().Where(book.IDEQ(id)).First(ctx)
	return book, err
}

func (r *bookRepository) UpdateBook(ctx context.Context, id int, title string, body string, userID int) (*ent.Book, error) {
	updatedBook, err := r.client.Book.
		UpdateOneID(id).
		SetTitle(title).
		SetBody(body).
		SetUserID(userID).
		Save(ctx)
	return updatedBook, err
}

func (r *bookRepository) DeleteBook(ctx context.Context, id int) error {
	err := r.client.Book.DeleteOneID(id).Exec(ctx)
	return err
}
