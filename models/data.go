package models

type Data struct {
	Coin_id  string `json:”id” bson:"id"`
	Rank     string `json:”rank” bson:"”rank”"`
	Symbol   string `json:”symbol” bson:”symbol”`
	Interval int    `json:”interval” bson:”interval”`
}
