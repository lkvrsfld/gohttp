package main

import "net/http"

// registers all handlers to api.multiplexer

func (api *Api)InitMultiplexer() (error) {
	api.multiplexer = http.NewServeMux()

	api.multiplexer.Handle("/", api.middleware.Handle(http.HandlerFunc(okHandler)))
	api.multiplexer.Handle("/err", api.middleware.Handle(http.HandlerFunc(errorHandler)))

	return nil
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello!"))
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write([]byte("Hello! Unavailable"))
}
