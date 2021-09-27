package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	VersionENV = "Vesion"
)

func greet(w *EngineResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func healthz(w *EngineResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func headers(w *EngineResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	if os.Getenv(VersionENV) == "" {
		os.Setenv(VersionENV, "1.0.0")

	}

	engine := &Engine{
		router: make(map[string]handler),
	}

	engine.AddRouter("/", greet)
	engine.AddRouter("/healthz", healthz)
	engine.AddRouter("/headers", headers)
	log.Fatal(http.ListenAndServe(":8080", engine))

}
