package main

import (
	"os"
	"log"
	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
)

type Decorator func(fasthttp.RequestHandler) fasthttp.RequestHandler

func Decorate(handler fasthttp.RequestHandler, decorators ...Decorator) fasthttp.RequestHandler {
	for _, decorator := range decorators {
		handler = decorator(handler)
	}

	return handler
}

func WithFunnyHeader(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		handler(ctx)
		ctx.Response.Header.Add("Funny-Header", "Ciao!")
	}
}

func WithAdditionalText(text string) Decorator {
	return func(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			handler(ctx)
			fmt.Fprintln(ctx, text)
		}
	}
}

func namedFunction(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("name"))
}

func main() {
	mux := fasthttprouter.New()
	logger := log.New(os.Stdout, "[server] ", log.LstdFlags|log.Lshortfile)

	server := &fasthttp.Server{
		Logger: logger,
		Handler: mux.Handler,
		Name: "server",
	}

	mux.GET("/", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintln(ctx, "Hello, world!")
	})

	mux.GET("/foo/*name", namedFunction)
	mux.GET("/boo/*name",
		Decorate(namedFunction,
			WithFunnyHeader,
			WithAdditionalText("Ciao!"),
		),
	)

	logger.Fatalln(
		server.ListenAndServe(":8080"),
	)
}
