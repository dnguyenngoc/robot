/* 
	@Author: Duy Nguyen
	@Email: <duynguyenngoc@hotmail.com>
*/

package database

import (
	"context"
	"time"
	"log"
	"fmt"
	"strconv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/dnguyenngoc/robot/service/settings"
)

const (
	connectTimeout = 10
	connectionStringTemplate = "mongodb://%s:%s@%s:%s/%s" 
)

func DBinstance() *mongo.Client  {
	/*
		Function for handle connect to mongodb with time out (10s) and make sure disable connect to db when occur incident
	*/

	conf := settings.GetEnv()
	
	mongoURI:= fmt.Sprintf(
		connectionStringTemplate, conf.Database.User, conf.Database.Pass, conf.Database.Host, 
		strconv.Itoa(conf.Database.Port), conf.Database.Database,
	)

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	defer cancel()

	// connect db with time out
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error while connecting to mongo: %v\n", err)
		log.Fatal(err)
	}

	// handle disconnect db
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Ping database
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return client
}

var Client *mongo.Client = DBinstance()



