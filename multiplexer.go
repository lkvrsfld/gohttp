package main

import "net/http"

// registers all handlers to api.multiplexer
func InitMultiplexer() (*http.ServeMux, error) {
	multiplexer := http.NewServeMux()

	multiplexer.Handle("/", api.middleware.Handle(http.HandlerFunc(okHandler)))
	multiplexer.Handle("/err", api.middleware.Handle(http.HandlerFunc(errorHandler)))
	return multiplexer, nil
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello!"))
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write([]byte("Hello! Unavailable"))
}
