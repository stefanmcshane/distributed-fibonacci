package main

import (
	"log"
	"net/http"

	"github.com/stefanmcshane/distributed-fibonacci/app/business"
)

func main() {

	// Connect to storage cache and
	// storage := business.NewInMemoryStorage()

	// Connect to database cache
	cfg := business.Config{
		User:       "postgres",
		Password:   "postgres",
		Host:       "localhost",
		Port:       5432,
		Name:       "postgres",
		DisableTLS: true,
	}
	storage, err := business.NewDatabaseConnection(cfg)
	if err != nil {
		log.Fatalf("Unable to connect to database with username:%s host:%s database:%s", cfg.User, cfg.Host, cfg.Name)
	}

	// setup fibonacci client
	rt := business.NewRTClient(storage)
	fib := NewFibApp(rt)

	mux := http.NewServeMux()
	mux.HandleFunc(endpointFibonacci, fib.handleFibonacci)
	mux.HandleFunc(endpointResultsUnder, fib.handleResultsUnder)
	mux.HandleFunc(endpointClearCache, fib.handleClearCache)

	middlewareMux := NewTimer(mux)
	log.Fatal(http.ListenAndServe(":8080", middlewareMux))
}
