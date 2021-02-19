package book_interfaces

import (
	"context"
	"github.com/booky/server/pkg/book_entities"
)

type BookGateway interface {
	GetBooks(ctx context.Context) ([]book_entities.Book, error)
	GetFilteredBooks(ctx context.Context, bookQuery book_entities.BookQuery) ([]book_entities.Book, error)
	UpsertBook(ctx context.Context, update book_entities.BookUpdate) error
	DeleteBook(ctx context.Context, isbn string) error
	ListenForChanges(ctx context.Context, c chan *book_entities.BookEvent) error
	GetEvents(ctx context.Context, count int64) ([]book_entities.BookEvent, error)
}
