package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Baraha/crypto_server.git/internal/config"
	"github.com/Baraha/crypto_server.git/pkg/adapters/db"
	"github.com/Baraha/crypto_server.git/pkg/models"
	"github.com/antonholmquist/jason"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

var coincapUrl = "https://api.coincap.io/v2/assets/"

func GetServerInfo(coin_id string) (response models.Data) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(coincapUrl + coin_id)
	req.Header.Set("Accept-Encoding", "gzip")
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// Perform the request
	for {
		err := fasthttp.Do(req, resp)
		if err != nil {
			fmt.Printf("Client get failed: %s\n", err)
			return
		}

		if resp.StatusCode() != fasthttp.StatusOK {
			fmt.Printf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
		}
		if resp.StatusCode() == fasthttp.StatusOK {
			fmt.Printf("Get data:  status code %d\n", resp.StatusCode())
			break
		}
	}
	contentType := resp.Header.Peek("Content-Type")
	if bytes.Index(contentType, []byte("application/json")) != 0 {
		fmt.Printf("Expected content type application/json but got %s\n", contentType)
		return
	}
	contentEncoding := resp.Header.Peek("Content-Encoding")
	var body []byte
	if bytes.EqualFold(contentEncoding, []byte("gzip")) {
		fmt.Println("Unzipping...")
		body, _ = resp.BodyGunzip()
	} else {
		body = resp.Body()
	}

	fmt.Printf("Response body is: %s", body)

	object, err := jason.NewObjectFromBytes(body)
	if err != nil {

		log.Fatal(err)
	}
	data_coin, err := object.GetObject("data")
	if err != nil {

		log.Fatal(err)
	}
	byte_data, err := data_coin.Marshal()

	if err != nil {

		log.Fatal(err)
	}
	var data models.Data
	err2 := json.Unmarshal(byte_data, &data)
	if err2 != nil {

		log.Fatal(err2)
	}
	fmt.Println("data", data)

	fmt.Println("json.Unmarshal(body, &data)", json.Unmarshal(body, &data))
	fmt.Println("resp to inner data : ", data)

	return data
}

func CoinView(ctx *fasthttp.RequestCtx) {
	/**
		*This is a comment.
		*@api {get} http://localhost:8080/cryptocurrency/
		*@apiName Просмотр валюты
		*@apiGroup Криптовалюта
		*@apiDescription Просмотр Криптовалюты, обноваляется в базе каждый цикл интервала
		*@apiParam {string} coin_id Вид просматриваемой валюты
		*@apiParam {string} rank Позиция валюты на мировой криптобирже
		*@apiParam {string} symbol символическое обожначение
		*@apiParam {int} interval Интервалы между обновлением валюты в секундах
		*@apiParam {string} priceUsd цена валюты в переводе в USD
		*@apiSuccessExample {json} Success-Response:
	        [
				{
					"_id": "611a7074e1d840f625b58c92",
					"coin_id": "bitcoin",
					"interval": 30,
					"priceusd": "46439.5197486433777590",
					"rank": "1",
					"symbol": "BTC"
				},
				{
					"_id": "611a70e5e1d840f625b58f1e",
					"coin_id": "ethereum",
					"interval": 1,
					"priceusd": "3228.3628716937351608",
					"rank": "2",
					"symbol": "ETH"
				},
				{
					"_id": "611a71cad88a76ef6f00de5f",
					"coin_id": "bitcoin",
					"interval": 30,
					"priceusd": "46439.5197486433777590",
					"rank": "1",
					"symbol": "BTC"
				},
				{
					"_id": "611a71d2c9a9566d5c98fb02",
					"coin_id": "bitcoin",
					"interval": 30,
					"priceusd": "46439.5197486433777590",
					"rank": "1",
					"symbol": "BTC"
				}
			]
	*/
	filter := bson.M{}

	var results []bson.M
	cur := db.Find(config.DB_NAME, models.COLLECTION_CRYPTOCURRENCY, filter)
	defer cur.Close(context.Background())

	cur.All(context.Background(), &results)
	json, _ := json.Marshal(results)
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.AppendBody(json)
}

func CreateCoinView(ctx *fasthttp.RequestCtx) {
	/**

	Создание партии
	@api {POST} api/batches/ Создание крипты
	@apiName Создания статистики по криптовалюте
	@apiGroup Криптовалюта
	@apiDescription Создания статистики по криптовалюте

	@apiParam {string} coin_id Вид просматриваемой валюты
	@apiParam {int} interval Интервалы между обновлением валюты в секундах

	@apiParamExample {json} Request-Example:
	{

		"coin_id": "bitcoin",
		"interval": 30
	}

	@apiSuccessExample {json} Success-Response:
	HTTP/1.1 200 OK
	{
		"InsertedID": "611a8824e450d2183ab5f9a2"
	}

	@apiError (500 BAD REQUEST) {Object} errors List of errors

	@apiErrorExample ValidationErrors:
	{
		{
			"message": "Failed to decode JSON object: Expecting value: line 1 column 1 (char 0)"
		}
	}



	*/

	var currency models.Data
	json.Unmarshal([]byte(ctx.Request.Body()), &currency)
	currency.ObjectID = primitive.NewObjectID()
	currency.Save()
	response, _ := json.Marshal("success")
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.AppendBody(response)

}

func DeleteCoinView(ctx *fasthttp.RequestCtx) {
	/**
	Удаление мониторинга за криптовалютой по ID

	@api {delete} /api/batches/<batch_id>/ Удаление мониторинга за криптовалютой
	@apiName Удаление мониторинга за криптовалютой
	@apiGroup Криптовалюта
	@apiDescription Удаление мониторинга за криптовалютой по ID
	@apiError (404 NOT FOUND) ID обьекта некорректен
	@apiError (404 NOT FOUND) {string} errors.common Common message
	@apiSuccessExample {json} Success-Response:
	{
		objID: ObjectID("611a8824e450d2183ab5f9a2")
		UserValue: 611a8824e450d2183ab5f9a2
		{"DeletedCount":0}
	}
	*/

	objID, _ := primitive.ObjectIDFromHex(fmt.Sprint(ctx.UserValue("id")))
	filter := bson.M{"_id": objID}
	db.DeleteOne(config.DB_NAME, models.COLLECTION_CRYPTOCURRENCY, filter)

}

func CoinItemView(ctx *fasthttp.RequestCtx) {

	filter := bson.M{}

	cur := db.Find(config.DB_NAME, models.COLLECTION_CRYPTOCURRENCY, filter)
	fmt.Println("cur: ", cur)

	cnt := 0
	var results map[string]models.Data
	results = make(map[string]models.Data)
	fmt.Println("results: ", results)
	for cur.Next(context.TODO()) {
		item := "item" + fmt.Sprint(cnt)
		fmt.Println("item: ", item)
		cnt++
		var elem models.Data
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("elem.Coin_id: ", elem.Coin_id)
		data := GetServerInfo(elem.Coin_id)
		fmt.Println("data append", data)
		results[item] = data
		fmt.Println("results update ", data)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	json, _ := json.Marshal(results)
	ctx.Response.AppendBody(json)
	time.Sleep(2 * time.Millisecond)

}
