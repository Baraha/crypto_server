package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Baraha/crypto_server.git/api"
	"github.com/Baraha/crypto_server.git/models"
	"github.com/Baraha/crypto_server.git/utils"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var dbName = "test"
var THROTTLE = 0

const (
	EXEC_THROTTLE = 10
)

func save_detect(data models.Data, id string, interval int) {

	if id == "" {
		log.Fatal("error id detected, id: ", id)
		return
	}
	fmt.Println("data.Interval: ", interval)
	fmt.Println("time: ", interval)
	time.Sleep(time.Duration(interval) * time.Second)
	collection_statistic, err := utils.GetMongoDbCollection(dbName, "cryptocurrency")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("id detected, id: ", id)
	filter := bson.M{"coin_id": id}
	fmt.Println("filter: ", filter)

	update := bson.M{"$set": bson.M{"rank": data.Rank, "symbol": data.Symbol, "priceusd": data.PriceUsd}}
	fmt.Println("data to save ! ", data)
	res, err := collection_statistic.UpdateMany(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
	THROTTLE--
}

func Control() {

	fmt.Println("Start working control service")
	collection_currency, err := utils.GetMongoDbCollection(dbName, "cryptocurrency")
	if err != nil {
		log.Fatal(err)
	}
	options := options.Find()
	filter := bson.M{}
	for THROTTLE < EXEC_THROTTLE+1 {
		time.Sleep(1 * time.Second)
		if THROTTLE < 0 {
			THROTTLE = 0
		}
		if THROTTLE >= EXEC_THROTTLE-1 {
			time.Sleep(4 * time.Second)
			fmt.Println("wait THROTTLE ", THROTTLE)
			continue
		}

		cur, err := collection_currency.Find(context.TODO(), filter, options)
		fmt.Println("cur: ", cur)
		if err != nil {
			log.Fatal(err)
		}

		cnt := 0
		for cur.Next(context.TODO()) {
			if THROTTLE < EXEC_THROTTLE-1 {
				THROTTLE++
			}
			cnt++
			var elem models.Data
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("elem.Coin_id: ", elem.Coin_id)
			if elem.Coin_id == "" {
				fmt.Println("error id detected, id: ", elem.Coin_id)
				continue
			}
			data := api.GetServerInfo(elem.Coin_id)

			go save_detect(data, elem.Coin_id, elem.Interval)

			fmt.Println("THROTTLE ", THROTTLE)
			fmt.Println("data append", data)

			fmt.Println("results update ", data)
		}
	}
}
