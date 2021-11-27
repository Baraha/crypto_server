package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Baraha/crypto_server.git/internal/api"
	"github.com/Baraha/crypto_server.git/internal/utils"
	"github.com/Baraha/crypto_server.git/models"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var dbName = "test"
var THROTTLE = 0

const (
	EXEC_THROTTLE = 10
)

func save_detect(data models.Data, id string, interval int, baseId models.Data) {

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

	update := bson.M{"$set": bson.M{
		"rank":     data.Rank,
		"symbol":   data.Symbol,
		"priceusd": data.PriceUsd,
	}}
	fmt.Println("data to save ! ", data)
	res, err := collection_statistic.UpdateMany(context.Background(), filter, update)
	fmt.Println("update append", update)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)

	update_one := bson.M{"$set": bson.M{
		"isHandle": false}}
	filter_one := bson.M{"_id": baseId.ObjectID}
	fmt.Println("data.ObjectID.Hex() ", data.ObjectID.Hex())

	res_one, err := collection_statistic.UpdateOne(context.Background(), filter_one, update_one)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", res_one.MatchedCount, res_one.ModifiedCount)
}

func detectServerInfo(elem models.Data) {
	data := api.GetServerInfo(elem.Coin_id)
	save_detect(data, elem.Coin_id, elem.Interval, elem)
	fmt.Println("data append", data)
	fmt.Println("results update ", data)
}

func Control() {

	fmt.Println("Start working control service")
	for {
		time.Sleep(1 * time.Second)
		collection_currency, err := utils.GetMongoDbCollection(dbName, "cryptocurrency")
		var results []models.Data
		if err != nil {
			log.Fatal(err)
		}
		options := options.Find()
		filter := bson.M{}

		cur, err := collection_currency.Find(context.TODO(), filter, options)
		fmt.Println("cur: ", cur)
		if err != nil {
			log.Fatal(err)
		}

		for cur.Next(context.TODO()) {
			var elem models.Data

			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("elem.IsHandle: ", elem.IsHandle)

			results = append(results, elem)
		}
		println("result for find", results)
		cur.Close(context.TODO())

		for index, value := range results {
			println("Index", index)
			fmt.Println("elem.Coin_id: ", value.Coin_id)
			if value.Coin_id == "" {
				fmt.Println("error id detected, id: ", value.Coin_id)
				continue
			}
			if value.IsHandle == false {
				update_one := bson.M{"$set": bson.M{
					"isHandle": true}}
				filter_one := bson.M{"_id": value.ObjectID}
				fmt.Println("elem.ObjectID.Hex() ", value.ObjectID.Hex())

				res, err := collection_currency.UpdateOne(context.Background(), filter_one, update_one)
				if err != nil {
					log.Fatal(err)
					return
				}
				go detectServerInfo(value)
				fmt.Printf("Matched %v documents in handler and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
				if value.IsHandle == true {
					continue
				}
			}
			fmt.Println("ishandle : ", value.IsHandle)

		}
		cur.Close(context.TODO())
		fmt.Println("next iter")

	}

}
