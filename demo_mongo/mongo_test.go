package demo_mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	component_mongo "go_demo/component_mongo"
	"testing"
)

func TestGetDatabases(t *testing.T) {
	mongoClient := component_mongo.NewLocalMongoClient()
	databases, err := mongoClient.ListDatabases(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	fmt.Println(databases)
}
