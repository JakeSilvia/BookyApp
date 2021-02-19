package handlers

import (
	"github.com/booky/server/handlers/booky_context"
	event_getter_interactor "github.com/booky/server/pkg/book_interactors/event_interactors"
	"github.com/booky/server/pkg/book_interfaces"
	"github.com/gocraft/web"
	"log"
	"net/http"
	"strconv"
)

type EventsGetterHandler struct {
	BookGateway book_interfaces.BookGateway
}

func (bg *EventsGetterHandler) Get (c *booky_context.ServerContext, w web.ResponseWriter, r *web.Request) {
	count := r.URL.Query().Get("count")
	countInt, err := strconv.Atoi(count)
	if err != nil {
		log.Printf("error converting count to int: %v", err)
		c.ServeErrorf(http.StatusBadRequest, err)
		return
	}

	getter := &event_getter_interactor.EventsGetter{
		BookGateway: bg.BookGateway,
		Count: countInt,
	}

	events, err := getter.Run(c.Ctx)
	if err != nil {
		log.Printf("error getting books: %v", err)
		c.ServeErrorf(http.StatusBadRequest, err)
		return
	}

	c.ServeJson(http.StatusOK, events)

}

func NewEventsGetterHandler (gateway book_interfaces.BookGateway) *EventsGetterHandler {
	return &EventsGetterHandler{BookGateway: gateway}
}
