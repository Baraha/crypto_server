package api

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Baraha/crypto_server.git/models"
	"github.com/Baraha/crypto_server.git/utils"
	"github.com/valyala/fasthttp"
)

func SaveLastStat(ctx *fasthttp.RequestCtx) {
	collection, err := utils.GetMongoDbCollection(dbName, "last_stat")
	if err != nil {
		log.Fatal(err)
		ctx.Response.Header.SetStatusCode(500)
		return
	}

	var last_stat models.Data
	json.Unmarshal([]byte(ctx.Request.Body()), &last_stat)

	res, err := collection.InsertOne(context.Background(), last_stat)
	if err != nil {
		log.Fatal(err)
		ctx.Response.Header.SetStatusCode(500)
		return
	}

	response, _ := json.Marshal(res)
	ctx.Response.AppendBody(response)

}
