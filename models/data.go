package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Data struct {
	ObjectID primitive.ObjectID `json:"-" bson:"_id"`
	Coin_id  string             `json:”coin_id” bson:"coin_id"`
	Rank     string             `json:”rank” bson:"”rank”"`
	Symbol   string             `json:”symbol” bson:”symbol”`
	Interval int                `json:”interval” bson:”interval”`
	PriceUsd string             `json:”priceUsd” bson:”priceUsd”`
	IsHandle bool               `json:”isHandle” bson:”isHandle”`
}
