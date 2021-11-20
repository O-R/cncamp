package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hzhhong/gap"

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

func main() {
	if os.Getenv(VersionENV) == "" {
		os.Setenv(VersionENV, "1.0.0")

	}
	logger := logx.NewStdLogger(os.Stdout)
	cfg := configx.New(flagconf)

	if err := cfg.Load(); err != nil {
		panic(err)
	}
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

	srv.UseSimple(func(c *gap.Context) {
		c.ResponseWriter.Header().Add(VersionENV, os.Getenv(VersionENV))
	})
	srv.Use(gap.LoggerProcessor(), gap.RouterProcessor())

	srv.AddRouter("/", greet)
	srv.AddRouter("/healthz", healthz)
	srv.AddRouter("/readinesshealthz", healthz)
	srv.AddRouter("/livenesshealthz", healthz)
	srv.AddRouter("/startuphealthz", healthz)
	srv.AddRouter("/headers", headers)

	return srv
}
