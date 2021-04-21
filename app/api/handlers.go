package main

import (
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/stefanmcshane/distributed-fibonacci/app/business"
)

type Fib struct {
	Client business.ReserveTrustClient
}

func NewFibApp(client business.ReserveTrustClient) Fib {
	return Fib{
		Client: client,
	}
}

func (f Fib) handleFibonacci(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println(r.URL)
	ordString := strings.TrimPrefix(r.URL.Path, endpointFibonacci)
	ordinal, ok := big.NewInt(0).SetString(ordString, 10)
	if !ok {
		httpError(w, fmt.Sprintf("error: parsing ordinal from request - %s\n", ordString), http.StatusBadRequest)
		return
	}
	fib, err := f.Client.Fibonacci(ctx, *ordinal)
	if err != nil {
		httpError(w, fmt.Sprintf("error: unable to get fib for ordinal %d - %s\n", ordinal, err.Error()), http.StatusInternalServerError)
		return
	}
	o := fibBigInt(*ordinal)
	fi := fibBigInt(fib)
	d := fibonacciResponse{
		Ordinal:   &o,
		Fibonacci: &fi,
	}
	httpRespondJson(w, d, http.StatusOK)
}

func (f Fib) handleResultsUnder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println(r.URL)
	fibString := strings.TrimPrefix(r.URL.Path, endpointResultsUnder)
	fib, ok := big.NewInt(0).SetString(fibString, 10)
	if !ok {
		httpError(w, fmt.Sprintf("error: parsing fibonacci number from request - %s\n", fibString), http.StatusBadRequest)
		return
	}

	count, err := f.Client.ResultsUnder(ctx, *fib)
	if err != nil {
		httpError(w, fmt.Sprintf("error: unable to get fib count under %d - %s\n", fib, err.Error()), http.StatusInternalServerError)
		return
	}
	fi := fibBigInt(*fib)

	d := fibonacciResponse{
		Count:     count,
		Fibonacci: &fi,
	}
	httpRespondJson(w, d, http.StatusOK)
}

func (f Fib) handleClearCache(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println(r.URL)

	err := f.Client.Storage.Clear(ctx)
	if err != nil {
		httpError(w, fmt.Sprintf("error: unable to clear cache - %s\n", err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
