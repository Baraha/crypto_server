package main

import (
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const (
	hosts      = "ds026491.mongolab.com:26491"
	database   = "messagingdb"
	username   = "admin"
	password   = "youPassword"
	collection = "messages"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Welcome!")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("name"))
}

func main() {
	r := router.New()
	r.GET("/", Index)
	r.GET("/hello/{name}", Hello)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))

}
