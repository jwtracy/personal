package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/johnwtracy/personal/src/apiserver/greeter"
)

var (
	port = flag.Int("port", 8080, "port for the server")
)

func main() {
	flag.Parse()

	greeterService := greeter.NewServer(
		"John Tracy",
		"See you, space cowboy!",
	)

	mux := http.NewServeMux()

	mux.Handle(greeterService.PathPrefix(), greeterService)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)
}
