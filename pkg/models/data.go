package models

import (
	"github.com/Baraha/crypto_server.git/internal/config"
	"github.com/Baraha/crypto_server.git/pkg/adapters/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const COLLECTION_CRYPTOCURRENCY = "cryptocurrency"

type Data struct {
	ObjectID primitive.ObjectID `json:"-" bson:"_id"`
	Coin_id  string             `json:”coin_id” bson:"coin_id"`
	Rank     string             `json:”rank” bson:"”rank”"`
	Symbol   string             `json:”symbol” bson:”symbol”`
	Interval int                `json:”interval” bson:”interval”`
	PriceUsd string             `json:”priceUsd” bson:”priceUsd”`
	IsHandle bool               `json:”isHandle” bson:”isHandle”`
}

func (Data *Data) Update() {

	db.UpdateByID(config.DB_NAME, COLLECTION_CRYPTOCURRENCY, Data, Data.ObjectID)
}

func (Data *Data) Save() {
	Data.ObjectID = primitive.NewObjectID()
	db.InsertOne(config.DB_NAME, COLLECTION_CRYPTOCURRENCY, Data)
}
