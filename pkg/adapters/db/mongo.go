package db

import (
	"context"
	"fmt"
	"log"

	"github.com/Baraha/crypto_server.git/pkg/utils/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var db_client *mongo.Client

func SetClient(active_client *mongo.Client) {
	db_client = active_client
}

func GetMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	collection := db_client.Database(DbName).Collection(CollectionName)
	return collection, nil
}

func CountDocs(DbName string, CollectionName string, filter interface{}) int64 {
	collection, err := GetMongoDbCollection(DbName, CollectionName)
	service.CatchErr(err)
	options_count := options.Count()
	count, _ := collection.CountDocuments(context.TODO(), filter, options_count)
	log.Printf("count documents %v", count)
	return count
}

func Find(DbName string, CollectionName string, filter interface{}) *mongo.Cursor {
	collection, err := GetMongoDbCollection(DbName, CollectionName)
	options_find := options.Find()
	options_count := options.Count()
	cur, err := collection.Find(context.TODO(), filter, options_find)
	service.CatchErr(err)
	count, _ := collection.CountDocuments(context.TODO(), filter, options_count)
	log.Printf("count documents %v", count)
	return cur

}

func InsertOne(DbName string, CollectionName string, data interface{}) {
	collection := db_client.Database(DbName).Collection(CollectionName)
	result, err := collection.InsertOne(context.TODO(), data, options.InsertOne())

	service.CatchErr(err)
	log.Printf("result create object : %v\n", result)

}

func DeleteOne(DbName string, CollectionName string, filter interface{}) {

	collection := db_client.Database(DbName).Collection(CollectionName)
	opt := options.Delete()
	result, err := collection.DeleteOne(context.TODO(), filter, opt)
	service.CatchErr(err)

	log.Printf("result delete object : %v\n", result)

}

func UpdateByID(DbName string, CollectionName string, document interface{}, ID primitive.ObjectID) {
	collection := db_client.Database(DbName).Collection(CollectionName)
	update := bson.M{
		"$set": document,
	}
	// id, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": ID}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	service.CatchErr(err)
	log.Printf("result update object : %v\n", result)

}

func UpdateOne(DbName string, CollectionName string, update bson.M, filter interface{}) {
	collection := db_client.Database(DbName).Collection(CollectionName)
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	service.CatchErr(err)
	fmt.Printf("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
	log.Printf("result update object : %v\n", res)

}

func UpdateMany(DbName string, CollectionName string, document interface{}, filter bson.M) {
	collection := db_client.Database(DbName).Collection(CollectionName)
	update := bson.M{
		"$set": document,
	}
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	service.CatchErr(err)
	fmt.Printf("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
	log.Printf("result update many object : %v\n", res)

}
