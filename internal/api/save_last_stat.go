package api

import (
	"encoding/json"
	"fmt"

	"github.com/Baraha/crypto_server.git/pkg/models"
	"github.com/valyala/fasthttp"
)

func SaveLastStat(ctx *fasthttp.RequestCtx) {

	var last_stat models.Data
	json.Unmarshal([]byte(ctx.Request.Body()), &last_stat)

	last_stat.Save()

	response, _ := json.Marshal(fmt.Sprintf("id : %v", last_stat.ObjectID.Hex()))
	ctx.Response.AppendBody(response)

}
