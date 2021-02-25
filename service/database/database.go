package database

import (
	"context"
	"time"
	"log"
	"fmt"
	"strconv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/dnguyenngoc/robot/service/settings"
)

const (
	connectTimeout = 10
	connectionStringTemplate = "mongodb://%s:%s@%s/%s:%s"
)


// InitialConnection will create new connection to mongo db
func DBinstance() *mongo.Client  {

	conf := settings.GetEnv()
	
	fmt.Println(conf.MongoDb.User)

	mongoURI:= fmt.Sprintf(connectionStringTemplate, conf.MongoDb.User, conf.MongoDb.Password, conf.MongoDb.Host, conf.MongoDb.Database, strconv.Itoa(conf.MongoDb.Port))

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error while connecting to mongo: %v\n", err)
	}

	return client
}

var Client *mongo.Client = DBinstance()



