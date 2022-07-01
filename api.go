package main

import (
	"net/http"
)

type Api struct {
	Addr        string
	server      *http.Server
	multiplexer *http.ServeMux
	middleware  MiddlewareChain
}

func (api *Api) init() error {
	var err error
	// init middleware
	err = api.InitMiddleware()
	if err != nil {
		return err
	}

	// init handlers
	err = api.InitMultiplexer()
	if err != nil {
		return err
	}
	// init handlers
	api.server = &http.Server{
		Addr:    api.Addr,
		Handler: api.multiplexer,
	}

	return nil
}

// starts the server with listenAndServe
func (api *Api) Start() error {
	return api.server.ListenAndServe()

}
