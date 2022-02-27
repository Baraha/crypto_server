package core

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Baraha/crypto_server.git/internal/config"
	"github.com/Baraha/crypto_server.git/pkg/adapters/db"
	"github.com/Baraha/crypto_server.git/pkg/models"
	"gopkg.in/mgo.v2/bson"
)

func workWithDataResults(results []models.Data) {
	for index, value := range results {
		if value.IsHandle {
			continue
		}

		fmt.Printf("Index %v\n", index)
		fmt.Printf("elem.Coin_id: %v\n", value.Coin_id)
		if value.Coin_id == "" {
			fmt.Printf("error id detected, id: %v", value.Coin_id)
			continue
		}
		if !value.IsHandle {
			update_one := bson.M{
				"isHandle": true}
			db.UpdateByID(config.DB_NAME, models.COLLECTION_CRYPTOCURRENCY, update_one, value.ObjectID)
			fmt.Printf("update data info %v", value)
			go DetectServerInfo(value)

		}
		fmt.Println("ishandle : ", value.IsHandle)

	}
	fmt.Println("next iter")

}

func DetectDataControl() {
	for {
		time.Sleep(1 * time.Second)
		var results []models.Data
		filter := bson.M{}

		cur := db.Find(config.DB_NAME, models.COLLECTION_CRYPTOCURRENCY, filter)

		fmt.Println("cur: ", cur)

		for cur.Next(context.TODO()) {
			var elem models.Data

			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("elem.IsHandle: ", elem.IsHandle)

			results = append(results, elem)
		}
		fmt.Println("result for find", results)

		workWithDataResults(results)
		cur.Close(context.TODO())

	}
}
