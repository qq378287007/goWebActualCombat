package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoCli *mongo.Client

func initDb() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	//连接到MongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func MgoCli() *mongo.Client {
	if mgoCli == nil {
		initDb()
	}
	return mgoCli
}
