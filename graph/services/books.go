package services

import (
	"context"
	"strconv"

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
		ID:   strconv.Itoa(inputBook.ID),
		Title: inputBook.Title,
		Genre: inputBook.Genre,
	}
}

func convertBookSlice(books models.BookSlice) []*model.Book {
	result := make([]*model.Book, 0, len(books))
	for _, book := range books {
		result = append(result, convertBook(book))
	}
	return result
}

func (u *bookService) GetBookByID(ctx context.Context, id int) (*model.Book, error) {
	book, err := models.FindBook(ctx, u.exec, id,
		models.BookColumns.ID, models.BookColumns.Title, models.BookColumns.Genre,
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
	).One(ctx, u.exec)
	if err != nil {
		return nil, err
	}
	return convertBook(book), nil
}


func (u *bookService) GetBookByGenre(ctx context.Context, genre string) (*model.Book, error) {
	book, err := models.Books(
		qm.Select(models.BookColumns.ID, models.BookColumns.Genre),
		models.BookWhere.Title.EQ(genre),
	).One(ctx, u.exec)
	if err != nil {
		return nil, err
	}
	return convertBook(book), nil
}

func (u *bookService) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
	books, err := models.Books().All(ctx, u.exec)
	if err != nil {
			return nil, err
	}
	return convertBookSlice(books), nil
}

func (u *bookService) CreateBook(ctx context.Context, input *model.NewBook) (*model.Book, error) {
	book := &models.Book{
		Title: input.Title,
		Genre: input.Genre,
	}
	err := book.Insert(ctx, u.exec, boil.Infer())
	if err != nil {
		return nil, err
	}
	return convertBook(book), nil
}