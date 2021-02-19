package book_listener_interactor

import (
	"context"
	"fmt"
	"github.com/booky/server/pkg/book_entities"
	"github.com/booky/server/pkg/book_interfaces"
	"github.com/gorilla/websocket"
)

type BooksListener struct {
	BookGateway book_interfaces.BookGateway
	Connection  *websocket.Conn
}

func (bg *BooksListener) Run(ctx context.Context) error {
	err := bg.validate()
	if err != nil {
		return err
	}

	messages := make(chan *book_entities.BookEvent, 0)
	errChan := make(chan error)
	go func() {
		err = bg.BookGateway.ListenForChanges(ctx, messages)
		if err != nil {
			errChan <- err
		}
	}()

	go func() {
		for {
			message := <-messages
			err = bg.Connection.WriteJSON(message)
			if err != nil {
				errChan <- err
			}
		}
	}()

	err = <-errChan
	return err
}

func (bg *BooksListener) validate() error {
	if bg.BookGateway == nil {
		return fmt.Errorf("missing book gateway")
	}
	if bg.Connection == nil {
		return fmt.Errorf("missing connection")
	}

	return nil
}
