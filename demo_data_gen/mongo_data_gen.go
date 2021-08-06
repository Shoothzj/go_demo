package demo_data_gen

import (
	"context"
	"demo_data-gen/game"
	"demo_data-gen/module"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	component_mongo "go_demo/component_mongo"
	component_util "go_demo/component_util"
	"time"
)

func GenData(mongoDataGenConfig *module.MongoDataGenConfig, count int) {
	mongoClient := component_mongo.NewMongoClient(mongoDataGenConfig.Host)
	collection := mongoClient.Database(mongoDataGenConfig.Database).Collection(mongoDataGenConfig.Collection)
	for i := 0; i < count; i++ {
		GenOneData(mongoDataGenConfig, collection)
	}
}

func GenOneData(mongoDataGenConfig *module.MongoDataGenConfig, collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	bsonData := bson.D{}
	for _, mongoFieldConfig := range mongoDataGenConfig.MongoFieldConfigs {
		filedType := mongoFieldConfig.FiledType
		fieldName := mongoFieldConfig.FieldName
		if filedType == module.Uuid {
			bsonData = append(bsonData, bson.E{Key: fieldName, Value: component_util.Uuid()})
		}
		if filedType == module.RandomGameName {
			bsonData = append(bsonData, bson.E{Key: fieldName, Value: RandomString(game.Game)})
		}
		if filedType == module.RandomDatetimeRange {
			panic("not supported yet.")
		}
		if filedType == module.RandomLongitude {
			bsonData = append(bsonData, bson.E{Key: fieldName, Value: RandomFloat64FromRange(73.40, 135.2)})
		}
		if filedType == module.RandomLatitude {
			bsonData = append(bsonData, bson.E{Key: fieldName, Value: RandomFloat64FromRange(3, 53)})
		}
	}
	_, err := collection.InsertOne(ctx, bsonData)
	if err != nil {
		panic(err)
	}
}
