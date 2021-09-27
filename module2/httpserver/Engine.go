package main

import (
	"log"
	"net/http"
	"os"
)

type handler func(*EngineResponseWriter, *http.Request)

type Engine struct {
	router map[string]handler
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Add(VersionENV, os.Getenv(VersionENV))

	w := &EngineResponseWriter{
		ResponseWriter: writer,
		StatusCode:     http.StatusOK,
	}
	if h, ok := e.router[req.URL.Path]; ok {
		h(w, req)
	}
	log.Printf("clientIp: %s; path: %s; statuscode: %d\n", req.RemoteAddr, req.URL.Path, w.StatusCode)
}

func (e *Engine) AddRouter(path string, h handler) {
	e.router[path] = h
}
