package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Baraha/crypto_server.git/internal/api"
	"github.com/Baraha/crypto_server.git/internal/services"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongo_url = "mongodb://172.17.0.3:27017/"

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

	r := router.New()
	r.GET("/cryptocurrency", api.CoinView)
	r.POST("/cryptocurrency", api.CreateCoinView)
	r.DELETE("/cryptocurrency/{id}", api.DeleteCoinView)
	r.GET("/cryptocurrency/analitics", api.CoinItemView)
	fmt.Printf("server starting")

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
	fmt.Println("server is start!")

}
