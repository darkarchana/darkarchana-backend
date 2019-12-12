package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/darkarchana/darkarchana-backend/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db string
var clientOptions *options.ClientOptions
var client *mongo.Client
var err error

// MongoDbSetup : setup Mongo DB
func MongoDbSetup() {
	// Set client options
	var uri string
	connectType := os.Getenv("DB_MONGO_CONNECT")
	if connectType == "local" {
		uri = fmt.Sprintf("mongodb://%s:%s",
			os.Getenv("DB_MONGO_LOCALHOST"),
			os.Getenv("DB_MONGO_PORT"),
		)
	} else if connectType == "cluster" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s?%s",
			os.Getenv("DB_MONGO_USERNAME"),
			os.Getenv("DB_MONGO_PASSWORD"),
			os.Getenv("DB_MONGO_CLUSTERHOST"),
			os.Getenv("DB_MONGO_OTHERARG"),
		)
	}

	log.Printf("Connecting to %s", uri)
	clientOptions = options.Client().ApplyURI(uri)

	if clientOptions == nil {
		log.Fatal("Database Not Found")
	} else {
		db = os.Getenv("DB_MONGO_DATABASE")
		log.Printf("INFO : Database That Will be Used is [%s]", db)
	}
}

// MongoDbConnect : Connect to Mongo DB
func MongoDbConnect() error {
	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Print(err)
		// Try Ping to MongoDB if Error in Connecting
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Print(err)
		}
		return err
	}
	log.Println("STATUS: Connected to MongoDB!")
	return nil
}

// MongoDbDisconnect : disconnect from Mongo DB
func MongoDbDisconnect() error {
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Print(err)
		return err
	}
	log.Println("STATUS: Connection to MongoDB closed.")
	return nil
}

// MongoDbFindOne : find one document
func MongoDbFindOne(dbOperate model.DbOperate) *mongo.SingleResult {
	// Select collection using dbOperate.Collection
	collection := client.Database(db).Collection(dbOperate.Collection)

	// Passing dbOperate.Filter as the filter matches one document in the collection
	result := collection.FindOne(context.TODO(), dbOperate.Filter)

	return result
}

// MongoDbFind : find all document
func MongoDbFind(dbOperate model.DbOperate) (*mongo.Cursor, error) {
	// Select collection using dbOperate.Collection
	collection := client.Database(db).Collection(dbOperate.Collection)

	// Pass these options from dbOperate.Option to the Find method
	findOptions := options.Find()
	if dbOperate.Option.FindLimit > 0 {
		findOptions.SetLimit(dbOperate.Option.FindLimit)
	}

	// Passing dbOperate.Filter as the filter matches all documents in the collection
	results, err := collection.Find(context.TODO(), dbOperate.Filter, findOptions)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return results, nil
}
