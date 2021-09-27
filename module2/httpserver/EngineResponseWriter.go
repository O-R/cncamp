package main

import "net/http"

type EngineResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *EngineResponseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
