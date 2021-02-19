package handlers

import (
	"github.com/booky/server/handlers/booky_context"
	"github.com/booky/server/pkg/book_interactors/book_getter_interactor"
	"github.com/booky/server/pkg/book_interfaces"
	"github.com/gocraft/web"
	"log"
	"net/http"
)

type BooksGetterHandler struct {
	BookGateway book_interfaces.BookGateway
}

func (bg *BooksGetterHandler) Get (c *booky_context.ServerContext, w web.ResponseWriter, r *web.Request) {
	getter := &book_getter_interactor.BooksGetter{
		BookGateway: bg.BookGateway,
	}

	books, err := getter.Run(c.Ctx)
	if err != nil {
		log.Printf("error getting books: %v", err)
		c.ServeErrorf(http.StatusBadRequest, err)
		return
	}

	c.ServeJson(http.StatusOK, books)

}

func NewBookGetterHandler (gateway book_interfaces.BookGateway) *BooksGetterHandler {
	return &BooksGetterHandler{BookGateway: gateway}
}
