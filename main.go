package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Baraha/crypto_server.git/api"
	"github.com/Baraha/crypto_server.git/models"
	"github.com/Baraha/crypto_server.git/services"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongo_url = "mongodb://mongo:27017/"
var dbName = "test"
var coincapUrl = "https://api.coincap.io/v2/assets/"

func main() {
	go services.Control()
	// Set client options
	clientOptions := options.Client().ApplyURI(mongo_url)

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

	bitcoin := models.Data{Coin_id: "bitcoin", Interval: 30}
	insertResult, err := collection.InsertOne(context.TODO(), bitcoin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	r := router.New()
	r.GET("/cryptocurrency", api.CoinView)
	r.POST("/cryptocurrency", api.CreateCoinView)
	r.DELETE("/cryptocurrency/{id}", api.DeleteCoinView)
	r.GET("/cryptocurrency/analitics", api.CoinItemView)
	fmt.Printf("server starting")

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
	fmt.Println("server is start!")

}
