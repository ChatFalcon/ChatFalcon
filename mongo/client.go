package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

// DB defines the MongoDB database.
var DB *mongo.Database

// ErrNoDocuments is used to echo the error from the MongoDB driver.
var ErrNoDocuments = mongo.ErrNoDocuments

// InitClient is used to initialise the client.
func InitClient() {
	MongoURI := os.Getenv("MONGO_URI")
	if MongoURI == "" {
		MongoURI = "mongodb://localhost:27017"
	}
	c, err := mongo.NewClient(options.Client().ApplyURI(MongoURI))
	if err != nil {
		panic(err)
	}
	err = c.Connect(context.TODO())
	if err != nil {
		panic(err)
	}
	DB = c.Database("chatfalcon")
}
