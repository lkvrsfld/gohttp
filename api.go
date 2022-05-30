package main

import (
	"fmt"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Api struct {
	Addr 		string 
	server 		*http.Server
	multiplexer *http.ServeMux
	middleware  []Middleware
}



func (api *Api) init() error {
	
	api.server = &http.Server{
		Addr:           api.Addr,
		Handler: 		api.multiplexer,
	}

	// init middleware
	middleware, err := api.InitMiddleware()
	if err != nil {
		return err
	}
	api.middleware = middleware

	// init handlers
	multiplexer, err := api.InitHandlers()
	if err != nil {
		return err
	}
	api.multiplexer = multiplexer

	return nil

}

// starts the server with listenAndServe
func (api *Api) Start() error {
	return api.server.ListenAndServe()
}

// registers all handlers to api.multiplexer
func (api *Api) InitHandlers() (*http.ServeMux,error) {
	multiplexer := http.NewServeMux()

	multiplexer.Handle("/", api.Handle(http.HandlerFunc(testHandler)))
	multiplexer.Handle("/download", api.Handle(http.HandlerFunc(testHandler)))

	return multiplexer, nil
}

// registers all middleware and writes to api.middlewares
func (api *Api) InitMiddleware() ([]Middleware, error) {
	var middleware []Middleware
	
	middleware = append(middleware,RateLimitMiddleware)

	return middleware, nil

}

// main handle function, iterates through all middleware from first to last, and handles the route

// solution: https://github.com/karankumarshreds/GoMiddlewares/blob/main/chain.go
func (api *Api) Handle(h http.Handler) http.Handler {
	var context = h
	for _, middleware := range api.middleware {
		fmt.Println("tete")
		context = middleware(context)
	}
	return context
}

//MIDDLEWARE
func RateLimitMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		// if !isValidRequest(ratelimit, ip) {
		// 	w.WriteHeader(http.StatusServiceUnavailable)
		// 	return
		// }

		
		// ratelimit.Hit(ip)
		fmt.Println("MIDDLEWARE")
		h.ServeHTTP(w, r)
	})
}

// func isValidRequest(l rl.Limit, key string) bool {
// 	_, ok := l.Rates[key]
// 	if !ok {
// 		return true
// 	}
// 	if l.Rates[key].Hits == l.MaxRequests {
// 		return false
// 	}
// 	return true
// }
