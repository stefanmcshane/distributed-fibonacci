package main

import (
	"fmt"
	"net/http"
	"time"
)

type Timer struct {
	handler http.Handler
}

//NewTimer constructs a new Timer middleware handler
func NewTimer(handler http.Handler) *Timer {
	return &Timer{handler}
}

func (t *Timer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	t.handler.ServeHTTP(w, r)
	duration := time.Now().Sub(startTime)
	fmt.Printf("%s served in %f seconds\n", r.URL, duration.Seconds())
}
