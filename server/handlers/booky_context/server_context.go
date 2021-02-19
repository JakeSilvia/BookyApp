package booky_context

import (
	"context"
	"encoding/json"
	"github.com/gocraft/web"
	"net/http"
)

type ServerContext struct {
	Ctx      context.Context
	w        web.ResponseWriter
	r        *web.Request
}

func (sc *ServerContext) InitServerContext(w web.ResponseWriter, r *web.Request, next web.NextMiddlewareFunc) {
	sc.Ctx = context.Background()
	sc.w = w
	sc.r = r
	next(w, r)
}

func (sc *ServerContext) ServeErrorf(status int, err error) {
	sc.ServeJson(status, err.Error())
}

func (sc *ServerContext) ServeJson(status int, jsonObject interface{}) {
	bts, err := json.Marshal(jsonObject)
	if err != nil {
		sc.w.WriteHeader(http.StatusInternalServerError)
		_, _ = sc.w.Write([]byte("unexpected error json encoding response"))
		return
	}

	sc.w.WriteHeader(status)
	_, _ = sc.w.Write(bts)
}
