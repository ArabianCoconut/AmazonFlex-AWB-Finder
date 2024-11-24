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


func ConnectAndUpload(awb string, datetime string, remark string) {
	var mongoConfig struct {
		Database     string
		DB_COLLECTION string
	}

	mongoConfig.Database = os.Getenv("DB_NAME")
	mongoConfig.DB_COLLECTION = os.Getenv("DB_COLLECTION")
	
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

func ConnectAndFetch() []bson.M {
	var mongoConfig struct {
		Database     string
		DB_COLLECTION string
	}

	mongoConfig.Database = os.Getenv("DB_NAME")
	mongoConfig.DB_COLLECTION = os.Getenv("DB_COLLECTION")
	
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_LOGIN"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println(err)
	}
	collection := client.Database(mongoConfig.Database).Collection(mongoConfig.DB_COLLECTION)
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
	}
	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		log.Println(err)
	}
	return results
}

func ConnectAndDelete(awb string) {
	var mongoConfig struct {
		Database      string
		DB_COLLECTION string
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("DB_LOGIN"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	collection := client.Database(mongoConfig.Database).Collection(mongoConfig.DB_COLLECTION)

	// Use FindOneAndDelete to delete a document and return the deleted one
	result := collection.FindOneAndDelete(context.Background(), bson.D{{Key: "awb", Value: awb}})
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No record found for AWB: %s", awb)
		} else {
			log.Printf("Error occurred while deleting AWB %s: %v", awb, err)
		}
		return
	}

	// Optional: Decode the deleted document if needed
	var deletedDoc bson.M
	if err := result.Decode(&deletedDoc); err != nil {
		log.Printf("Failed to decode deleted document for AWB %s: %v", awb, err)
	} else {
		log.Printf("Deleted document: %v", deletedDoc)
	}

	log.Printf("AWB %s deleted successfully", awb)
}