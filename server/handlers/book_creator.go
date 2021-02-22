package handlers

import (
	"encoding/json"
	"github.com/booky/server/handlers/booky_context"
	"github.com/booky/server/pkg/book_entities"
	"github.com/booky/server/pkg/book_interactors/book_creator_interactor"
	"github.com/booky/server/pkg/book_interfaces"
	"github.com/gocraft/web"
	"log"
	"net/http"
)

type BookCreatorHandler struct {
	BookGateway book_interfaces.BookGateway
}

func (bg *BookCreatorHandler) Put (c *booky_context.ServerContext, w web.ResponseWriter, r *web.Request) {
	book := book_entities.BookUpdate{}
	err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
		log.Printf("error decoding body: %v", err)
		c.ServeErrorf(http.StatusBadRequest, err)
		return
	}

	upserter := &book_creator_interactor.BookCreator{
		BookGateway: bg.BookGateway,
		Book: book,
	}
	err = upserter.Run(c.Ctx)
	if err != nil {
		log.Printf("error creating book: %v", err)
		c.ServeErrorf(http.StatusBadRequest, err)
		return
	}

	c.ServeJson(http.StatusOK, "ok")
}

func NewBookCreatorHandler (gateway book_interfaces.BookGateway) *BookCreatorHandler {
	return &BookCreatorHandler{BookGateway: gateway}
}
