package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Baraha/crypto_server.git/models"
	"github.com/Baraha/crypto_server.git/utils"
	"github.com/antonholmquist/jason"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var dbName = "test"
var coincapUrl = "https://api.coincap.io/v2/assets/"

func GetServerInfo(coin_id string) (response models.Data) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(coincapUrl + coin_id)
	req.Header.Set("Accept-Encoding", "gzip")
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// Perform the request
	err := fasthttp.Do(req, resp)
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)
		return
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		fmt.Printf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
		return
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
	data_coin, err := object.GetObject("data")
	byte_data, err3 := data_coin.Marshal()
	if err3 != nil {

		log.Fatal(err3)
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

	options := options.Find()
	//options.SetLimit(5)
	filter := bson.M{}
	collection, err := utils.GetMongoDbCollection(dbName, "cryptocurrency")
	if err != nil {
		log.Fatal(err)
		ctx.Response.Header.SetStatusCode(500)
		return
	}
	var results []bson.M
	cur, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		log.Fatal(err)
		ctx.Response.Header.SetStatusCode(500)
	}

	defer cur.Close(context.Background())

	if err != nil {
		log.Fatal(err)
		ctx.Response.Header.SetStatusCode(500)
		return
	}

	cur.All(context.Background(), &results)

	if results == nil {
		log.Fatal(err)
		ctx.Response.Header.SetStatusCode(500)
		return
	}

	json, _ := json.Marshal(results)
	ctx.Response.AppendBody(json)
}

func CreateCoinView(ctx *fasthttp.RequestCtx) {
	collection, err := utils.GetMongoDbCollection(dbName, "cryptocurrency")
	if err != nil {
		log.Fatal(err)
		ctx.Response.Header.SetStatusCode(500)
		return
	}

	var currency models.StatCoin
	json.Unmarshal([]byte(ctx.Request.Body()), &currency)

	res, err := collection.InsertOne(context.Background(), currency)
	if err != nil {
		log.Fatal(err)
		ctx.Response.Header.SetStatusCode(500)
		return
	}

	response, _ := json.Marshal(res)
	ctx.Response.AppendBody(response)

}

func DeleteCoinView(ctx *fasthttp.RequestCtx) {
	collection, err := utils.GetMongoDbCollection(dbName, "cryptocurrency")

	if err != nil {
		log.Fatal(err)
		ctx.Response.Header.SetStatusCode(500)
		return
	}
	objID, _ := primitive.ObjectIDFromHex(fmt.Sprint(ctx.UserValue("id")))
	fmt.Fprintf(ctx, "objID: %s\n", objID)
	fmt.Fprintf(ctx, "UserValue: %s\n", ctx.UserValue("id"))
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		log.Fatal(err)
		ctx.Response.Header.SetStatusCode(500)
		return
	}

	jsonResponse, _ := json.Marshal(res)
	ctx.Response.AppendBody(jsonResponse)
}

func CoinItemView(ctx *fasthttp.RequestCtx) {

	collection, err := utils.GetMongoDbCollection(dbName, "cryptocurrency")
	if err != nil {
		log.Fatal(err)
		ctx.Response.Header.SetStatusCode(500)
		return
	}

	options := options.Find()
	//options.SetLimit(2)
	filter := bson.M{}

	// Here's an array in which you can store the decoded documents

	// Passing nil as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), filter, options)
	fmt.Println("cur: ", cur)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	cnt := 0
	var results map[string]models.Data
	results = make(map[string]models.Data)
	fmt.Println("results: ", results)
	for cur.Next(context.TODO()) {
		time.Sleep(1 * time.Millisecond)
		item := "item" + fmt.Sprint(cnt)
		fmt.Println("item: ", item)
		cnt++
		// create a value into which the single document can be decoded
		var elem models.StatCoin
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
}
