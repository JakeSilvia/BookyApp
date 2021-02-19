package book_mock_gateway

import (
	"context"
	"github.com/booky/server/pkg/book_entities"
)

type BookMockGateway struct {
	Books []book_entities.Book
	Error error
}

func (bmg *BookMockGateway) GetBooks(ctx context.Context) ([]book_entities.Book, error) {
	return nil, nil
}

func (bmg *BookMockGateway) GetFilteredBooks(ctx context.Context, bookQuery book_entities.BookQuery) ([]book_entities.Book, error) {
	return nil, nil
}

func (bmg *BookMockGateway) UpsertBook(ctx context.Context, update book_entities.BookUpdate) error {
	return nil
}

func (bmg *BookMockGateway) DeleteBook(ctx context.Context, isbn string) error {
	return nil
}

func (bmg *BookMockGateway) ListenForChanges(ctx context.Context, c chan *book_entities.BookMessage) error {
	return nil
}
