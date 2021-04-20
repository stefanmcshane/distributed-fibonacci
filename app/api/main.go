package main

import (
	"fmt"
	"net/http"

	"github.com/stefanmcshane/distributed-fibonacci/app/business"
)

func main() {

	// Connect to storage cache and setup fibonacci client
	storage := business.NewInMemoryStorage()
	rt := business.NewRTClient(storage)
	fib := NewFibApp(rt)

	http.HandleFunc(fmt.Sprintf("%s", endpointFibonacci), fib.handleFibonacci)
	http.ListenAndServe(":8080", nil)
}
