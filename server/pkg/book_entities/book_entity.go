package book_entities

const (
	BookStatusCheckedOut = "Checked Out"
	BookStatusAvailable  = "Available"

	BookActionUpdated = "updated"
	BookActionCreated = "created"
	BookActionDeleted = "deleted"
)

type Book struct {
	Title       string `json:"title" bson:"title"`
	Author      string `json:"author" bson:"author"`
	ISBN        string `json:"isbn" bson:"isbn"`
	Description string `json:"description" bson:"description"`
	Status      string `json:"status" bson:"status"`
}

type BookEvent struct {
	Action   string `json:"action"`
	BookItem *Book  `json:"book_item"`
	Date     string `json:"date"`
}

type BookUpdate struct {
	Title       *string `json:"title"`
	Author      *string `json:"author"`
	ISBN        string  `json:"isbn"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

type BookQuery struct {
	Title       *string `json:"title"`
	Author      *string `json:"author"`
	ISBN        *string `json:"isbn"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

type BookMessage struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Action      string `json:"action"`
}

func (b *Book) ToList() []string {
	return []string{
		b.Title,
		b.Author,
		b.ISBN,
		b.Description,
		b.Status,
	}
}
