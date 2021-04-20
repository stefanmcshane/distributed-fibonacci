package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/stefanmcshane/distributed-fibonacci/app/business"
)

const (
	endpointFibonacci = "/fibonacci/"
)

type Fib struct {
	Client business.ReserveTrustClient
}

func NewFibApp(client business.ReserveTrustClient) Fib {
	return Fib{
		Client: client,
	}
}

type fibonacciResponse struct {
	Ordinal   int64 `json:"ordinal"`
	Fibonacci int64 `json:"fibonacci"`
}

func (f Fib) handleFibonacci(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println(r.URL)
	ordString := strings.TrimPrefix(r.URL.Path, endpointFibonacci)
	ordinal, err := strconv.ParseInt(ordString, 10, 64)
	if err != nil {
		httpError(w, fmt.Sprintf("error: parsing ordinal from request - %s\n", ordString), http.StatusBadRequest)
		return
	}
	fib, err := f.Client.Fibonacci(ctx, ordinal)
	if err != nil {
		httpError(w, fmt.Sprintf("error: unable to get fib for ordinal %d - %s\n", ordinal, err.Error()), http.StatusInternalServerError)
		return
	}
	d := fibonacciResponse{
		Ordinal:   ordinal,
		Fibonacci: fib,
	}
	httpRespondJson(w, d, http.StatusOK)
}

func httpError(w http.ResponseWriter, message string, statusCode int) {
	fmt.Println(message)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
	w.Header().Add("Content-Type", "text/plain")

}

func httpRespondJson(w http.ResponseWriter, data interface{}, statusCode int) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		httpError(w, fmt.Sprintf("Error returning json request - %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
}
