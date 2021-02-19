package book_deleter_interactor

import (
	"context"
	"fmt"
	"github.com/booky/server/pkg/book_interfaces"
)

type BookDeleter struct {
	BookGateway book_interfaces.BookGateway
	ISBN string
}

func (bd *BookDeleter) Run(ctx context.Context) error {
	if err := bd.validate(); err != nil {
		return err
	}

	return bd.BookGateway.DeleteBook(ctx, bd.ISBN)
}

func (bd *BookDeleter) validate () error {
	if bd.BookGateway == nil {
		return fmt.Errorf("missing book gateway")
	}

	if bd.ISBN == "" {
		return fmt.Errorf("missing ISBN")
	}

	return nil
}
