package main

import (
	"04/model"
	"04/mongodb"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var (
		client     = mongodb.MgoCli()
		collection *mongo.Collection
		err        error
		uResult    *mongo.UpdateResult
	)
	collection = client.Database("my_db").Collection("table1")
	filter := bson.M{"jobName": "job multi1"}
	update := bson.M{"$set": model.
		UpdateByJobName{Command: "byModel", Content: "model"}}
	if uResult, err = collection.
		UpdateMany(context.TODO(), filter, update); err != nil {
		log.Fatal(err)
	}
	//uResult.MatchedCount表示符合过滤条件的记录数，即更新了多少条数据。
	log.Println(uResult.MatchedCount)
}
