package main

import (
	"fmt"
	"log"

	"github.com/fasthttp/router"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(w, "hello, %s!\n", ctx.UserValue("name"))
}

func main() {
	r := router.New()
	r.GET("/", Index)
	r.GET("/hello/{name}", Hello)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
