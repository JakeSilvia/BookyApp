package book_creator_interactor
import (
	"context"
	"fmt"
	"github.com/booky/server/pkg/book_entities"
	"github.com/booky/server/pkg/book_interfaces"
)

type BookCreator struct {
	BookGateway book_interfaces.BookGateway
	Book        book_entities.BookUpdate
}

func (bu *BookCreator) Run(ctx context.Context) error {
	if err := bu.validate(); err != nil {
		return err
	}

	query := book_entities.BookQuery{
		ISBN: &bu.Book.ISBN,
	}

	book, err := bu.BookGateway.GetFilteredBooks(ctx, query)
	if err != nil {
		return err
	}

	if len(book) > 0 {
		return fmt.Errorf("ISBN [%v] already exists", bu.Book.ISBN)
	}

	return bu.BookGateway.UpsertBook(ctx, bu.Book)
}

func (bu *BookCreator) validate() error {
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
