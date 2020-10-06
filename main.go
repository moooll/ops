package main

import (
	"opsd/svc"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fasthttp/router"
	"github.com/gobuffalo/pop/v5"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func main() {
	{
		globalLogger, _ := zap.NewDevelopment()
		zap.ReplaceGlobals(globalLogger)
	}

	r := router.New()

	db, err := pop.Connect(os.Getenv("DB_ENV"))
	if err != nil {
		zap.L().Fatal("could not connect to database", zap.Error(err))
	}
	i := svc.New(db)
	i.RouterExtender(r.Group("/v1"))

	server := &fasthttp.Server{
		Handler: r.Handler,
	}

	go func() {
		zap.L().Info("Starting HTTP server on port 8888")
		err := server.ListenAndServe(":8888")
		if err != nil {
			zap.L().With(zap.Error(err)).Fatal("could not start server")
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	zap.L().Info("Stopping Server...")
	select {
	case <-time.After(15 * time.Second):
	default:
		err = server.Shutdown()
		if err != nil {
			zap.L().Error("could not shutdown server", zap.Error(err))
		}
	}

	zap.L().Info("Server Stopped")
}
