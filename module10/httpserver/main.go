package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/hzhhong/cncamp/httpserver/metrics"
	"github.com/hzhhong/gap"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	configx "github.com/hzhhong/gap/config"
	logx "github.com/hzhhong/gap/log"
)

const (
	VersionENV = "Vesion"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "conf", "configs/config.yaml", "config path, eg: -conf config.yaml")
}

func greet(ctx *gap.Context) {
	fmt.Fprintf(ctx.ResponseWriter, "Hello World! %s", time.Now())
}

func healthz(ctx *gap.Context) {
	ctx.ResponseWriter.WriteHeader(http.StatusOK)
}

func headers(ctx *gap.Context) {
	for name, headers := range ctx.Request.Header {
		for _, h := range headers {
			fmt.Fprintf(ctx.ResponseWriter, "%v: %v\n", name, h)
		}
	}
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func hello(ctx *gap.Context) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	user := ctx.Request.URL.Query().Get("user")
	delay := randInt(0, 2000)
	time.Sleep(time.Millisecond * time.Duration(delay))
	if user != "" {
		io.WriteString(ctx.ResponseWriter, fmt.Sprintf("hello [%s], delay %d\n", user, delay))
	} else {
		io.WriteString(ctx.ResponseWriter, fmt.Sprintf("hello [stranger], delay %d\n", delay))
	}
	io.WriteString(ctx.ResponseWriter, "===================Details of the http request header:============\n")
	for k, v := range ctx.Request.Header {
		io.WriteString(ctx.ResponseWriter, fmt.Sprintf("%s=%s\n", k, v))
	}
}

func main() {
	if os.Getenv(VersionENV) == "" {
		os.Setenv(VersionENV, "1.0.0")

	}
	logger := logx.NewStdLogger(os.Stdout)
	cfg := configx.New(flagconf)

	if err := cfg.Load(); err != nil {
		panic(err)
	}
	metrics.Register()
	srv1 := newHttpServer(
		getCfgValue(cfg, "server.http1.name", logger),
		getCfgValue(cfg, "server.http1.addr", logger),
		logger)
	srv2 := newHttpServer(
		getCfgValue(cfg, "server.http2.name", logger),
		getCfgValue(cfg, "server.http2.addr", logger),
		logger)
	app := gap.NewApp(gap.Servers(srv1, srv2))

	log.Print("App Started")
	if err := app.Run(); err != nil {
		panic(err)
	}
	log.Print("App Exited Properly")

}

func getCfgValue(cfg configx.Config, key string, logger logx.Logger) string {
	if val, err := cfg.Value(key); err == nil {
		if valstring, ok := val.(string); ok {

			logx.With(logger,
				"method", "getCfgValue",
			).Log(logx.LevelInfo, key, valstring)
			return valstring
		} else {
			panic(fmt.Errorf("get config value error, key: %s", key))
		}
	} else {
		panic(err)
	}
}

func newHttpServer(name string, addr string, logger logx.Logger) *gap.Server {
	srv := gap.RawSrv(name, addr, logger)

	useIngoreMiddleware(srv, "/favicon.ico")
	usePromMiddleware(srv, "/metrics")
	useHealthzMiddleware(srv, "/healthz")
	useHealthzMiddleware(srv, "/readinesshealthz")
	useHealthzMiddleware(srv, "/livenesshealthz")
	useHealthzMiddleware(srv, "/startuphealthz")

	srv.UseSimple(func(c *gap.Context) {
		c.ResponseWriter.Header().Add(VersionENV, os.Getenv(VersionENV))
	})
	srv.Use(gap.LoggerProcessor(), gap.RouterProcessor())

	srv.AddRouter("/", greet)
	srv.AddRouter("/headers", headers)
	srv.AddRouter("/hello", hello)
	return srv
}

func useHealthzMiddleware(srv *gap.Server, path string) {

	middleware := func(next gap.MiddlewareHandler) gap.MiddlewareHandler {
		return func(c *gap.Context) {

			if c.Request.URL.Path == path {
				healthz(c)
			} else if next != nil {
				next(c)
			}
		}
	}
	srv.HttpHandler.Use(middleware)
}

func useIngoreMiddleware(srv *gap.Server, path string) {

	middleware := func(next gap.MiddlewareHandler) gap.MiddlewareHandler {
		return func(c *gap.Context) {

			if c.Request.URL.Path == path {

			} else if next != nil {
				next(c)
			}
		}
	}
	srv.HttpHandler.Use(middleware)
}

func usePromMiddleware(srv *gap.Server, path string) {
	promhandler := promhttp.Handler()
	prommiddleware := func(next gap.MiddlewareHandler) gap.MiddlewareHandler {
		return func(c *gap.Context) {

			if c.Request.URL.Path == path {
				promhandler.ServeHTTP(c.ResponseWriter.ResponseWriter, c.Request)
			} else if next != nil {
				next(c)
			}
		}
	}
	srv.HttpHandler.Use(prommiddleware)
}
