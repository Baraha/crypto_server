package app

import (
	"context"
	"fmt"
	"log"

	"github.com/Baraha/crypto_server.git/internal/api"
	"github.com/Baraha/crypto_server.git/internal/config"
	"github.com/Baraha/crypto_server.git/internal/currency_tracker"
	"github.com/Baraha/crypto_server.git/pkg/adapters/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Start() {
	config.Init("shell")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.MONGO_URL))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	db.SetClient(client)

	defer client.Disconnect(context.Background())

	go currency_tracker.Control()

	// Set client options
	fmt.Println("server is starting!")
	api.Init()
	api.Start()

}
