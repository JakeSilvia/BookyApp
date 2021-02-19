package handlers

import (
	"github.com/booky/server/handlers/booky_context"
	"github.com/booky/server/pkg/book_interactors/book_listener_interactor"
	"github.com/booky/server/pkg/book_interfaces"
	"github.com/gocraft/web"
	"github.com/gorilla/websocket"
	"net/http"
)

type BooksWebsocketHandler struct {
	BookGateway book_interfaces.BookGateway
}

var upgrader = websocket.Upgrader{	CheckOrigin: func(r *http.Request) bool {
	return true
},
}
func (bws *BooksWebsocketHandler) Get (c *booky_context.ServerContext, w web.ResponseWriter, r *web.Request) {
	conn, err := upgrader.Upgrade(w, r.Request, nil)
	if err != nil {
		c.ServeErrorf(http.StatusBadRequest, err)
		return
	}
	defer conn.Close()
	listener := book_listener_interactor.BooksListener{
		BookGateway: bws.BookGateway,
		Connection: conn,
	}
	err = listener.Run(c.Ctx)
	if err != nil {
		c.ServeErrorf(http.StatusBadRequest, err)
		return
	}

	c.ServeJson(http.StatusOK, "ok")
}

func NewBooksWebsocketHandler (gateway book_interfaces.BookGateway) *BooksWebsocketHandler {
	return &BooksWebsocketHandler{BookGateway: gateway}
}
