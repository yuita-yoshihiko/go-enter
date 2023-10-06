package services

import (
	"context"
	"log"
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

func (u *bookService) ListBooksByID(ctx context.Context, IDs []int) ([]*model.Book, error) {
	books, err := models.Books(
		qm.Select(models.BookColumns.ID, models.BookColumns.Title),
		models.BookWhere.ID.IN(IDs),
	).All(ctx, u.exec)
	if err != nil {
		return nil, err
	}
	return convertBookSlice(books), nil
}

func (u *bookService) CreateBook(ctx context.Context, input *model.NewBook) (*model.Book, error) {
	// ログ1: 関数の開始
	log.Println("Creating a book with title:", input.Title)

	// データベースに新しい本を追加する試み
	book := &models.Book{
		Title: input.Title,
	}

	// ログ2: インサート前
	log.Println("Attempting to insert the book into the DB.")

	// データベースに本をインサート
	err := book.Insert(ctx, u.exec, boil.Infer())
	if err != nil {
		// ログ3: インサートエラー
		log.Printf("Error inserting the book into the DB: %v", err)
		return nil, err
	}

	// ログ4: インサート成功
	log.Println("Successfully inserted the book into the DB.")

	// 成功した場合、変換して返す
	return convertBook(book), nil
}