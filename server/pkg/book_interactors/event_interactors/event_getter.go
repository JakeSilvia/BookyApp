package event_getter_interactor

import (
	"context"
	"fmt"
	"github.com/booky/server/pkg/book_entities"
	"github.com/booky/server/pkg/book_interfaces"
)

type EventsGetter struct {
	BookGateway book_interfaces.BookGateway
	Count int
}

func (bg *EventsGetter) Run(ctx context.Context) ([]book_entities.BookEvent, error) {
	if err := bg.validate(); err != nil {
		return nil, err
	}

	return bg.BookGateway.GetEvents(ctx, int64(bg.Count))
}

func (bg *EventsGetter) validate () error {
	if bg.BookGateway == nil {
		return fmt.Errorf("missing book gateway")
	}

	return nil
}
