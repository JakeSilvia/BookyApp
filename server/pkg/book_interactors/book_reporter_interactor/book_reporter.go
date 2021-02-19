package book_reporter_interactor

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/booky/server/pkg/book_entities"
	"github.com/booky/server/pkg/book_interfaces"
	"reflect"
)

type BooksReporter struct {
	BookGateway book_interfaces.BookGateway
}

func (br *BooksReporter) Run(ctx context.Context) ([]byte, error) {
	if err := br.validate(); err != nil {
		return nil, err
	}

	books, err := br.BookGateway.GetBooks(ctx)
	if err != nil {
		return nil, err
	}

	reportBts, err := parseBooksToCsv(books)
	if err != nil {
		return nil, err
	}

	return reportBts.Bytes(), err
}

func (br *BooksReporter) validate () error {
	if br.BookGateway == nil {
		return fmt.Errorf("missing book gateway")
	}

	return nil
}

func parseBooksToCsv(books []book_entities.Book) (*bytes.Buffer, error) {
	bookList := [][]string{getBookFields()}
	for _, book := range books {
		bookList = append(bookList, book.ToList())
	}

	b := new(bytes.Buffer)
	writer := csv.NewWriter(b)
	err := writer.WriteAll(bookList)
	return b, err
}

func getBookFields() []string {
	fields := make([]string, 0)
	val := reflect.ValueOf(book_entities.Book{})
	for i := 0; i < val.Type().NumField(); i++ {
		fields = append(fields, val.Type().Field(i).Name)
	}

	return fields
}


