// es saber que url va a que handler

package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Printf("Hello World %s", w)
}
