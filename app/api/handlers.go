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
	d := fibonacciResponse{
		Ordinal:   fibBigInt(*ordinal),
		Fibonacci: fibBigInt(fib),
	}
	httpRespondJson(w, d, http.StatusOK)
}
