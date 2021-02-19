package main

import (
	"context"
	"fmt"
	"github.com/booky/server/data_gateways/mongo"
	"github.com/booky/server/handlers"
	"github.com/booky/server/handlers/booky_context"
	"github.com/booky/server/pkg/book_interfaces"
	"github.com/booky/server/settings"
	"github.com/gocraft/web"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func Health (c *booky_context.ServerContext, w web.ResponseWriter, r *web.Request) {
	c.ServeJson(http.StatusOK, "ok")
}

func makeBookDataGateway(ctx context.Context) (book_interfaces.BookGateway, error) {
	mongoOptions := options.Client().ApplyURI(settings.MongoConnection)
	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		return nil, fmt.Errorf("error parsing connection string: %s", err)
	}

	bng := &book_mongo_gateway.BookMongoGateway {
		DB: client.Database("test"),
	}
	return bng, nil
}

func main () {
	ctx := context.Background()
	gateway, err := makeBookDataGateway(ctx)
	if err != nil {
		log.Fatalf("error making gateway: %v", err)
	}
	router := web.NewWithPrefix(booky_context.ServerContext{}, "/api").
		Middleware((*booky_context.ServerContext).InitServerContext).
		Get("/books", handlers.NewBookGetterHandler(gateway).Get).
		Post("/books", handlers.NewBookUpdateHandler(gateway).Post).
		Put("/books", handlers.NewBookCreatorHandler(gateway).Put).
		Delete("/books", handlers.NewBookDeleteHandler(gateway).Delete).
		Get("/books/report", handlers.NewBookReportHandler(gateway).Get).
		Get("/events", handlers.NewEventsGetterHandler(gateway).Get).
		Get("/health", Health).
		Get("/connect", handlers.NewBooksWebsocketHandler(gateway).Get)


	fs := http.FileServer(http.Dir("./static"))
	mux := http.NewServeMux()
	mux.Handle("/api/", router)
	mux.Handle("/", fs)

	log.Printf("listening on: %v", ":" + settings.Port)
	port := "5000"
	if settings.Port != "" {
		port = settings.Port
	}
	err = http.ListenAndServe(":" + port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
