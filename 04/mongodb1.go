package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//1.建立连接
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017").SetConnectTimeout(5*time.Second))
	if err != nil {
		fmt.Print(err)
		return
	}

	//2.选择数据库 my_db
	db := client.Database("my_db")

	//3.选择表 my_collection
	collection := db.Collection("my_collection")
	fmt.Print(collection)
}
