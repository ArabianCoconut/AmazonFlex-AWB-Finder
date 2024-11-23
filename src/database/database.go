package database

import (
	"context"
	"log"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectandUpload connects to a MongoDB database and uploads a document with the provided details.
// It retrieves the database configuration from environment variables.
//
// Parameters:
//   - awb: The airway bill number to be inserted.
//   - datetime: The datetime string to be inserted.
//   - remark: The remark string to be inserted.
//
// Environment Variables:
//   - DB_NAME: The name of the database.
//   - DB_COLLECTION: The name of the collection.
//   - DB_LOGIN: The MongoDB connection URI.
//
// Logs:
//   - Logs an error if the connection to the database fails.
//   - Logs an error if the document insertion fails.
//   - Logs a success message if the document is inserted successfully.
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
