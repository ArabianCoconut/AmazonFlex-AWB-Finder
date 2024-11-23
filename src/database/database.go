package database

import (
	"context"
	"log"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connectandupload connects to a MongoDB database and uploads a document.
// It first attempts to load environment variables from a local .env file.
// If the .env file is not found, it falls back to loading environment variables from the system.
// The function then uses these environment variables to configure the MongoDB client and collection.
// Finally, it inserts a document with a single field "name" and value "pi" into the specified collection.
func ConnectandUpload(awb string, datetime string, remark string) {
	var mongoConfig struct {
		Database     string
		DB_COLLECTION string
		DB_LOGIN      string
	}

	mongoConfig.Database = os.Getenv("DB_NAME")
	mongoConfig.DB_COLLECTION = os.Getenv("DB_COLLECTION")
	// mongoConfig.DB_LOGIN = os.Getenv("DB_LOGIN")

	
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_LOGIN"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println(err)
	}
	collection := client.Database(mongoConfig.Database).Collection(mongoConfig.DB_COLLECTION)
	_, err = collection.InsertOne(context.Background(), bson.D{
		{Key: "awb", Value: awb},
		{Key: "datetime", Value: datetime},
		{Key: "remark", Value: remark},
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Document inserted successfully")
	}
}
