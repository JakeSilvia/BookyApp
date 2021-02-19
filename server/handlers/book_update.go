package handlers

import (
	"encoding/json"
	"github.com/booky/server/handlers/booky_context"
	"github.com/booky/server/pkg/book_entities"
	"github.com/booky/server/pkg/book_interactors/book_upsert_interactor"
	"github.com/booky/server/pkg/book_interfaces"
	"github.com/gocraft/web"
	"log"
	"net/http"
)

type BooksUpdateHandler struct {
	BookGateway book_interfaces.BookGateway
}

func (bg *BooksUpdateHandler) Post (c *booky_context.ServerContext, w web.ResponseWriter, r *web.Request) {
	book := book_entities.BookUpdate{}
	err := json.NewDecoder(r.Body).Decode(&book)
	upserter := &book_upsert_interactor.BookUpserter{
		BookGateway: bg.BookGateway,
		Book: book,
	}
	err = upserter.Run(c.Ctx)
	if err != nil {
		log.Printf("error updating book: %v", err)
		c.ServeErrorf(http.StatusBadRequest, err)
		return
	}

	c.ServeJson(http.StatusOK, "ok")
}

func NewBookUpdateHandler (gateway book_interfaces.BookGateway) *BooksUpdateHandler {
	return &BooksUpdateHandler{BookGateway: gateway}
}
