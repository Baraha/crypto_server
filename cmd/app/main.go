package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Baraha/crypto_server.git/internal/api"
	"github.com/Baraha/crypto_server.git/internal/services"
	"github.com/Baraha/crypto_server.git/internal/utils"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	go services.Control()
	// Set client options
	fmt.Println("server is starting!")
	clientOptions := options.Client().ApplyURI(utils.Mongo_url)

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

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))

}
