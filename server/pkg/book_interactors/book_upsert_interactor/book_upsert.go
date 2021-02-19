package book_upsert_interactor

import (
	"context"
	"fmt"
	"github.com/booky/server/pkg/book_entities"
	"github.com/booky/server/pkg/book_interfaces"
)

type BookUpserter struct {
	BookGateway book_interfaces.BookGateway
	Book        book_entities.BookUpdate
}

func (bu *BookUpserter) Run(ctx context.Context) error {
	if err := bu.validate(); err != nil {
		return err
	}

	return bu.BookGateway.UpsertBook(ctx, bu.Book)
}

func (bu *BookUpserter) validate() error {
	if bu.BookGateway == nil {
		return fmt.Errorf("missing book gateway")
	}

	if bu.Book.ISBN == "" {
		return fmt.Errorf("IBSN is a required field")
	}

	if bu.Book.Status != nil {
		switch *bu.Book.Status {
		case book_entities.BookStatusAvailable:
			fallthrough
		case book_entities.BookStatusCheckedOut:
			break
		default:
			return fmt.Errorf("invalid status: %v", *bu.Book.Status)
		}
	}
	return nil
}
