package main

import (
	"fmt"
	"net/http"
)

type Middleware func(http.Handler) http.Handler
type MiddlewareChain []Middleware

// registers all middleware and writes to api.middlewares
func InitMiddleware() (MiddlewareChain, error) {
	var middlewareChain MiddlewareChain

	middlewareChain = append(middlewareChain, RateLimitMiddleware)

	return middlewareChain, nil
}

func RateLimitMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("MIDDLEWARE")
		h.ServeHTTP(w, r)
	})
}

// main handle function, iterates through all middleware from first to last, and handles the route

func (mc MiddlewareChain) Handle(originalHandler http.Handler) http.Handler {
	if originalHandler == nil {
		originalHandler = http.DefaultServeMux
	}

	for i := range mc {
		originalHandler = mc[len(mc)-1-i](originalHandler)
	}
	return originalHandler
}