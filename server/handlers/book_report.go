package handlers

import (
	"fmt"
	"github.com/booky/server/handlers/booky_context"
	"github.com/booky/server/pkg/book_interactors/book_reporter_interactor"
	"github.com/booky/server/pkg/book_interfaces"
	"github.com/gocraft/web"
	"log"
	"net/http"
	"time"
)

type BooksReportHandler struct {
	BookGateway book_interfaces.BookGateway
}

func (bg *BooksReportHandler) Get (c *booky_context.ServerContext, w web.ResponseWriter, r *web.Request) {
	getter := &book_reporter_interactor.BooksReporter{
		BookGateway: bg.BookGateway,
	}

	report, err := getter.Run(c.Ctx)
	if err != nil {
		log.Printf("error getting books report: %v", err)
		c.ServeErrorf(http.StatusBadRequest, err)
		return
	}

	filename := fmt.Sprintf("attachment; filename=%v-books.csv", time.Now().String())
	w.Header().Set("Content-Disposition", filename)
	w.WriteHeader(http.StatusOK)
	w.Write(report)
}

func NewBookReportHandler (gateway book_interfaces.BookGateway) *BooksReportHandler {
	return &BooksReportHandler{BookGateway: gateway}
}
