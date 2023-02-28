package main

import (
	"fmt"
	"net/http"
	"time"
)

type Api struct {
	Host        string
	Port        string
	server      *http.Server
	multiplexer *http.ServeMux
	middleware  MiddlewareChain
}

func (api *Api) Init() error {
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
		Addr:    api.Host + ":" + api.Port,
		Handler: api.multiplexer,
	}

	return nil
}

// starts the server with listenAndServe
func (api *Api) Start() error {
	return api.server.ListenAndServe()
}

func (api *Api) Log(log string) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	prefix := timestamp + ": "
	fmt.Println(prefix + log)
}
