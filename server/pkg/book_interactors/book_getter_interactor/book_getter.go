package book_getter_interactor

import (
	"context"
	"fmt"
	"github.com/booky/server/pkg/book_entities"
	"github.com/booky/server/pkg/book_interfaces"
)

type BooksGetter struct {
	BookGateway book_interfaces.BookGateway
}

func (bg *BooksGetter) Run(ctx context.Context) ([]book_entities.Book, error) {
	if err := bg.validate(); err != nil {
		return nil, err
	}

	return bg.BookGateway.GetBooks(ctx)
}

func (bg *BooksGetter) validate () error {
	if bg.BookGateway == nil {
		return fmt.Errorf("missing book gateway")
	}

	return nil
}
