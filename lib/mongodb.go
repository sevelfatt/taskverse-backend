package lib

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var MongoClient *mongo.Client

func ConnectMongoDB(){
	connectionString := os.Getenv("MONGODB_URI")
	if connectionString == "" {
		log.Fatal("MONGODB_URI not set")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

  // Ping the primary
    if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
        panic(err)
    }
    log.Println("Successfully connected to MongoDB!")

	MongoClient = client
}