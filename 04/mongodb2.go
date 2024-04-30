package main

import (
	"gitee.com/shirdonl/goWebActualCombat/chapter4/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var (
		client     = mongodb.MgoCli()
		db         *mongo.Database
		collection *mongo.Collection
	)
	//选择数据库 my_db
	db = client.Database("my_db")

	//选择表 my_collection
	collection = db.Collection("my_collection")
	collection = collection
}
