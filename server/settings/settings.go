package settings

import "os"

var (
	MongoConnection = os.Getenv("MONGO")
	Port = os.Getenv("PORT")
)
