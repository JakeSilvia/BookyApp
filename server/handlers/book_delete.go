package handlers

import (
	"github.com/booky/server/handlers/booky_context"
	"github.com/booky/server/pkg/book_interactors/book_deleter_interactor"
	"github.com/booky/server/pkg/book_interfaces"
	"github.com/gocraft/web"
	"log"
	"net/http"
)

type BooksDeleteHandler struct {
	BookGateway book_interfaces.BookGateway
}

func (bg *BooksDeleteHandler) Delete (c *booky_context.ServerContext, w web.ResponseWriter, r *web.Request) {
	query := r.URL.Query()
	isbn := query.Get("isbn")
	deleter := &book_deleter_interactor.BookDeleter{
		BookGateway: bg.BookGateway,
		ISBN: isbn,
	}
	err := deleter.Run(c.Ctx)
	if err != nil {
		log.Printf("error deleting book [%s]: %v", isbn, err)
		c.ServeErrorf(http.StatusBadRequest, err)
		return
	}

	c.ServeJson(http.StatusOK, "ok")
}

func NewBookDeleteHandler (gateway book_interfaces.BookGateway) *BooksDeleteHandler {
	return &BooksDeleteHandler{BookGateway: gateway}
}
