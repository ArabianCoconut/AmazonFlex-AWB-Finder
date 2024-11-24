// ConnectAndUpload connects to the MongoDB database and uploads a document with the given AWB, datetime, and remark.
// It retrieves the database connection details from environment variables DB_LOGIN, DB_NAME, and DB_COLLECTION.
// If the document is inserted successfully, it logs a success message; otherwise, it logs the error.
//
// Parameters:
//   - awb: The Air Waybill number to be uploaded.
//   - datetime: The datetime associated with the AWB.
//   - remark: Any remark associated with the AWB.

// ConnectAndFetch connects to the MongoDB database and fetches all documents from the specified collection.
// It retrieves the database connection details from environment variables DB_LOGIN, DB_NAME, and DB_COLLECTION.
// It returns a slice of bson.M containing all the documents in the collection.
//
// Returns:
//   - []bson.M: A slice of bson.M containing all the documents in the collection.

// ConnectAndDelete connects to the MongoDB database and deletes a document with the given AWB.
// It retrieves the database connection details from environment variables DB_LOGIN, DB_NAME, and DB_COLLECTION.
// If the document is deleted successfully, it logs a success message; otherwise, it logs the error.
// If no document is found with the given AWB, it logs a message indicating that no record was found.
//
// Parameters:
//   - awb: The Air Waybill number to be deleted.
package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)




func ConnectAndUpload(awb string, datetime string, remark string) {
	
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_LOGIN"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println(err)
	}
	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLLECTION"))
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

	clientOptions := options.Client().ApplyURI(os.Getenv("DB_LOGIN"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println(err)
	}
	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLLECTION"))
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

	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLLECTION"))

	result := collection.FindOneAndDelete(context.Background(), bson.D{{Key: "awb", Value: awb}})
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No record found for AWB: %s", awb)
		} else {
			log.Printf("Error occurred while deleting AWB %s: %v", awb, err)
		}
		return
	}

	var deletedDoc bson.M
	if err := result.Decode(&deletedDoc); err != nil {
		log.Printf("Failed to decode deleted document for AWB %s: %v", awb, err)
	} else {
		log.Printf("Deleted document: %v", deletedDoc)
	}

	log.Printf("AWB %s deleted successfully", awb)
}