package services

import (
	"context"

	"go-enter/graph/model"

	"github.com/volatiletech/sqlboiler/boil"
)

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/service_mock.go
type Services interface {
	BookService
}

type BookService interface {
	GetBookByID(ctx context.Context, id string) (*model.Book, error)
	GetBookByTitle(ctx context.Context, title string) (*model.Book, error)
	ListBooksByID(ctx context.Context, IDs []string) ([]*model.Book, error)
}


type services struct {
	*bookService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		bookService:        &bookService{exec: exec},
	}
}