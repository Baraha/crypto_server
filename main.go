package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Welcome!")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("name"))
}

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017/")

	// Connect to MongoDB
	Client, _ := mongo.Connect(context.TODO(), clientOptions)

	// Check the connection
	err := Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	r := router.New()
	r.GET("/", Index)
	r.GET("/hello/{name}", Hello)
	fmt.Printf("server starting")

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
	fmt.Println("server is start!")

}
