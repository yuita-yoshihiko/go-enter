package services

import (
	"context"

	"go-enter/src/database/models"
	"go-enter/graph/model"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type bookService struct {
	exec boil.ContextExecutor
}

func convertBook(inputBook *models.Book) *model.Book {
	return &model.Book{
		ID:   inputBook.ID,
		Title: inputBook.Title,
	}
}

func convertBookSlice(books models.BookSlice) []*model.Book {
	result := make([]*model.Book, 0, len(books))
	for _, book := range books {
		result = append(result, convertBook(book))
	}
	return result
}

func (u *bookService) GetBookByID(ctx context.Context, id string) (*model.Book, error) {
	book, err := models.FindBook(ctx, u.exec, id,
		models.BookColumns.ID, models.BookColumns.Title,
	)
	if err != nil {
		return nil, err
	}
	return convertBook(book), nil
}

func (u *bookService) GetBookByTitle(ctx context.Context, title string) (*model.Book, error) {
	book, err := models.Books(
		qm.Select(models.BookColumns.ID, models.BookColumns.Title),
		models.BookWhere.Title.EQ(title),
		// qm.Where("name = ?", name),
	).One(ctx, u.exec)
	if err != nil {
		return nil, err
	}
	return convertBook(book), nil
}

func (u *bookService) ListBooksByID(ctx context.Context, IDs []string) ([]*model.Book, error) {
	books, err := models.Books(
		qm.Select(models.BookColumns.ID, models.BookColumns.Title),
		models.BookWhere.ID.IN(IDs),
	).All(ctx, u.exec)
	if err != nil {
		return nil, err
	}
	return convertBookSlice(books), nil
}