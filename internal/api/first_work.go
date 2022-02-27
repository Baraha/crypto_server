package api

import (
	"fmt"
	"log"

	"github.com/Baraha/crypto_server.git/internal/config"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var rout *router.Router

func Init() {
	r := router.New()
	r.GET("/cryptocurrency", CoinView)
	r.POST("/cryptocurrency", CreateCoinView)
	r.DELETE("/cryptocurrency/{id}", DeleteCoinView)
	r.GET("/cryptocurrency/analitics", CoinItemView)
	rout = r
}

func Start() {
	log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf(":%v", config.SERVICE_PORT), rout.Handler))
}
