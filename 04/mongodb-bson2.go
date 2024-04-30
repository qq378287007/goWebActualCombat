package main

import (
	"04/mongodb"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var (
		client     = mongodb.MgoCli()
		collection *mongo.Collection
		err        error
		cursor     *mongo.Cursor
	)
	collection = client.Database("my_db").Collection("table1")
	//按照jobName分组,countJob中存储每组的数目
	groupStage := mongo.Pipeline{bson.D{
		{"$group", bson.D{
			{"_id", "$jobName"},
			{"countJob", bson.D{
				{"$sum", 1},
			}},
		}},
	}}
	if cursor, err = collection.Aggregate(context.TODO(), groupStage); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
