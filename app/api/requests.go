package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
)

const (
	endpointFibonacci = "/fibonacci/"
)

type fibonacciResponse struct {
	Ordinal   fibBigInt `json:"ordinal"`
	Fibonacci fibBigInt `json:"fibonacci"`
}

type fibBigInt big.Int

func (fbi fibBigInt) MarshalJSON() ([]byte, error) {
	i2 := big.Int(fbi)
	return []byte(fmt.Sprintf(`"%s"`, i2.String())), nil
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
