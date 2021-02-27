/* 
	@Author: Duy Nguyen
	@Email: <duynguyenngoc@hotmail.com>
*/

package database

import (
	"context"
	"time"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	connectTimeout = 60
	connectionStringTemplate = "mongodb://%s:%s@%s:%s/%s"
	mongoURI = "mongodb://robot:1q2w3e4r@127.0.0.1:27017/robot" // need handle prod env
	databaseName = "robot" // need handle prod env
)

var Client *mongo.Client = DBinstance()

func DBinstance() *mongo.Client  {
	/*
		Function for handle connect to mongodb with time out (10s) and make sure disable connect to db when occur incident
	*/

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal(err)
    }

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	defer cancel()

	// connect db with time out
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error while connecting to mongo: %v\n", err)
	}

	// Ping database
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return client
}


//OpenCollection is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	
    var collection *mongo.Collection = client.Database("robot").Collection(collectionName)

    return collection
}

