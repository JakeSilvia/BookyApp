package book_mongo_gateway

import (
	"context"
	"github.com/booky/server/pkg/book_entities"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	BooksCollection = "books"
	EventsCollection = "events"
)

type BookMongoGateway struct {
	DB          *mongo.Database
	BookChannel chan book_entities.Book
}

type Change struct {
	FullDocument  *book_entities.BookEvent `bson:"fullDocument"`
	OperationType string `bson:"operationType"`
	DocumentKey   struct {
		Id primitive.ObjectID `bson:"_id"`
	} `bson:"documentKey"`
}

func (bmg *BookMongoGateway) GetEvents(ctx context.Context, count int64) ([]book_entities.BookEvent, error) {
	opts := &options.FindOptions{Limit: &count}
	cursor, err := bmg.DB.Collection(EventsCollection).Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	var events []book_entities.BookEvent
	err = cursor.All(ctx, &events)
	return events, err
}

func (bmg *BookMongoGateway) GetBooks(ctx context.Context) ([]book_entities.Book, error) {
	cursor, err := bmg.DB.Collection(BooksCollection).Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, err
	}

	var books []book_entities.Book
	err = cursor.All(ctx, &books)
	return books, err
}

func (bmg *BookMongoGateway) GetFilteredBooks(ctx context.Context, bookQuery book_entities.BookQuery) ([]book_entities.Book, error) {
	query := parseQuery(bookQuery)
	cursor, err := bmg.DB.Collection(BooksCollection).Find(ctx, query, nil)
	if err != nil {
		return nil, err
	}

	var books []book_entities.Book
	err = cursor.All(ctx, &books)
	return books, err
}

func parseQuery(q book_entities.BookQuery) bson.M {
	query := bson.M{}
	if q.Title != nil {
		query["Title"] = *q.Title
	}
	if q.Author != nil {
		query["Author"] = *q.Author
	}
	if q.Description != nil {
		query["Description"] = *q.Description
	}
	if q.Status != nil {
		query["Status"] = *q.Status
	}
	if q.ISBN != nil {
		query["ISBN"] = *q.ISBN
	}

	return query
}

func (bmg *BookMongoGateway) UpsertBook(ctx context.Context, update book_entities.BookUpdate) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"$set": parseUpdate(update)}
	_, err := bmg.DB.Collection(BooksCollection).UpdateOne(ctx, bson.M{"ISBN": update.ISBN}, updateQuery, opts)
	result := bmg.DB.Collection(BooksCollection).FindOne(ctx, bson.M{"ISBN": update.ISBN})
	bookUpdated := &book_entities.Book{}
	err = result.Decode(bookUpdated)
	if err != nil {
		return err
	}

	return bmg.storeEvent(ctx, book_entities.BookActionUpdated, bookUpdated)
}

func parseUpdate(update book_entities.BookUpdate) bson.M {
	query := bson.M{}
	if update.Title != nil {
		query["Title"] = *update.Title
	}
	if update.Author != nil {
		query["Author"] = *update.Author
	}
	if update.Description != nil {
		query["Description"] = *update.Description
	}
	if update.Status != nil {
		query["Status"] = *update.Status
	}

	return query
}

func (bmg *BookMongoGateway) DeleteBook(ctx context.Context, isbn string) error {
	result := bmg.DB.Collection(BooksCollection).FindOne(ctx, bson.M{"ISBN": isbn})
	_, err := bmg.DB.Collection(BooksCollection).DeleteOne(ctx, bson.M{"ISBN": isbn})
	if err != nil {
		return err
	}

	if result == nil {
		return nil
	}

	bookToDelete := &book_entities.Book{}
	err = result.Decode(bookToDelete)
	if err != nil {
		return err
	}

	return bmg.storeEvent(ctx, book_entities.BookActionDeleted, bookToDelete)
}

func (bmg *BookMongoGateway) storeEvent(ctx context.Context, action string, book *book_entities.Book) error {
	_, err := bmg.DB.Collection(EventsCollection).InsertOne(ctx, book_entities.BookEvent{
		Action: action,
		BookItem: book,
		Date: time.Now().String(),
	})
	return err
}

func (bmg *BookMongoGateway) ListenForChanges(ctx context.Context, c chan *book_entities.BookEvent) error {
	updateLookup := options.UpdateLookup
	episodesStream, err := bmg.DB.Collection(EventsCollection).Watch(context.TODO(), mongo.Pipeline{}, &options.ChangeStreamOptions{FullDocument: &updateLookup})
	if err != nil {
		panic(err)
	}

	for {
		change, err := iterateChangeStream(ctx, episodesStream)
		if err != nil {
			return err
		}

		c <- change.FullDocument
	}
}

func iterateChangeStream(routineCtx context.Context, stream *mongo.ChangeStream) (*Change, error) {
	for stream.Next(routineCtx) {
		change := &Change{}
		if err := stream.Decode(change); err != nil {
			return nil, err
		}

		return change, nil
	}
	return nil, nil
}
