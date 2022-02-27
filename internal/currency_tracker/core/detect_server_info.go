package core

import (
	"fmt"
	"log"
	"time"

	"github.com/Baraha/crypto_server.git/internal/api"
	"github.com/Baraha/crypto_server.git/internal/config"
	"github.com/Baraha/crypto_server.git/pkg/adapters/db"
	"github.com/Baraha/crypto_server.git/pkg/models"
	"gopkg.in/mgo.v2/bson"
)

func DetectServerInfo(elem models.Data) {
	data := api.GetServerInfo(elem.Coin_id)
	Save_detect(data, elem.Coin_id, elem.Interval, elem)
	fmt.Println("data append", data)
	fmt.Println("results update ", data)
}

func Save_detect(data models.Data, id string, interval int, baseId models.Data) {

	if id == "" {
		log.Fatal("error id detected, id: ", id)
		return
	}
	fmt.Println("data.Interval: ", interval)
	fmt.Println("time: ", interval)
	time.Sleep(time.Duration(interval) * time.Second)

	filter := bson.M{"coin_id": id}
	fmt.Println("filter: ", filter)

	update := bson.M{
		"rank":     data.Rank,
		"symbol":   data.Symbol,
		"priceusd": data.PriceUsd,
	}

	db.UpdateMany(config.DB_NAME, models.COLLECTION_CRYPTOCURRENCY, update, filter)
	fmt.Println("update append", update)

	update_one := bson.M{
		"isHandle": false}

	db.UpdateByID(config.DB_NAME, models.COLLECTION_CRYPTOCURRENCY, update_one, baseId.ObjectID)
	fmt.Println("data.ObjectID.Hex() ", data.ObjectID.Hex())

}
