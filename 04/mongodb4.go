package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"04/model"
	"04/mongodb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var (
		client     = mongodb.MgoCli()
		err        error
		collection *mongo.Collection
		result     *mongo.InsertManyResult
		id         primitive.ObjectID
	)
	collection = client.Database("my_db").Collection("table1")

	//批量插入
	result, err = collection.InsertMany(context.TODO(), []interface{}{
		model.LogRecord{
			JobName: "job multi1",
			Command: "echo multi1",
			Err:     "",
			Content: "1",
			Tp: model.ExecTime{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		model.LogRecord{
			JobName: "job multi2",
			Command: "echo multi2",
			Err:     "",
			Content: "2",
			Tp: model.ExecTime{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if result == nil {
		log.Fatal("result nil")
	}
	for _, v := range result.InsertedIDs {
		id = v.(primitive.ObjectID)
		fmt.Println("自增ID", id.Hex())
	}
}
