package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func panicOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s :%s", msg, err)
	}
}

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// Instance
var MI MongoInstance

func SetupDatabase() {
	uri := os.Getenv("MONGODB_URI")

	var err error
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	panicOnError(err, "Failed to new client")
	log.Println("New client success")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	panicOnError(err, "Failed to connect database")
	log.Println("Database connected")

	databaseName := os.Getenv("MONGODB_DATABASE")
	MI = MongoInstance{
		Client: client,
		DB:     client.Database(databaseName),
	}
}
