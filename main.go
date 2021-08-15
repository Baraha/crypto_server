package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

type StatCoin struct {
	Coin_id  string `json:”coin_id”`
	Interval int    `json:”interval”`
}

var mongo_uri = "mongodb://mongo:27017/"
var dbName = "test"

func GetMongoDbConnection() (*mongo.Client, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongo_uri))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client, nil
}

func getMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	client, err := GetMongoDbConnection()

	if err != nil {
		return nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)

	return collection, nil
}

func CoinView(ctx *fasthttp.RequestCtx) {
	options := options.Find()
	//options.SetLimit(5)
	filter := bson.M{}
	collection, err := getMongoDbCollection(dbName, "cryptocurrency")
	if err != nil {
		log.Fatal(err)
		return
	}
	var results []bson.M
	cur, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.Background())

	if err != nil {
		log.Fatal(err)
		return
	}

	cur.All(context.Background(), &results)

	if results == nil {
		log.Fatal(err)
		return
	}

	json, _ := json.Marshal(results)
	ctx.Response.AppendBody(json)
}

func CreateCoin(ctx *fasthttp.RequestCtx) {
}
func DeleteCoin(ctx *fasthttp.RequestCtx) {

}

func CoinItemView(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("name"))
}

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI(mongo_uri)

	// Connect to MongoDB
	Client, _ := mongo.Connect(context.TODO(), clientOptions)

	// Check the connection
	err := Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// add object in data
	collection := Client.Database("test").Collection("cryptocurrency")

	bitcoint := StatCoin{"bitcoin", 2}
	insertResult, err := collection.InsertOne(context.TODO(), bitcoint)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	r := router.New()
	r.GET("/", CoinView)
	r.GET("/hello/{name}", CoinItemView)
	fmt.Printf("server starting")

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
	fmt.Println("server is start!")

}
