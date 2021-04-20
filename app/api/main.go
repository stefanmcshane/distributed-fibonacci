package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stefanmcshane/distributed-fibonacci/app/business"
)

func main() {

	// Connect to storage cache and setup fibonacci client
	storage := business.NewInMemoryStorage()
	rt := business.NewRTClient(storage)
	fib := NewFibApp(rt)

	mux := http.NewServeMux()
	mux.HandleFunc(fmt.Sprintf("%s", endpointFibonacci), fib.handleFibonacci)

	middlewareMux := NewTimer(mux)
	log.Fatal(http.ListenAndServe(":8080", middlewareMux))
}
