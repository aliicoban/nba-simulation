package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/alicobanserver/internal/service"
	httptransport "github.com/alicobanserver/internal/transport/http"
)

const (
	httpServerAddr         = ":4444"
	httpServerReadTimeout  = 10 * time.Second
	httpServerWriteTimeout = 10 * time.Second
)

func main() {
	ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelFunc()

	var s service.Service
	{
		s = service.NewService()
	}

	var hs http.Server
	{
		hh := httptransport.MakeHTTPHandler(ctx, s)

		hs = http.Server{
			Addr:         httpServerAddr,
			ReadTimeout:  httpServerReadTimeout,
			WriteTimeout: httpServerWriteTimeout,
			Handler:      hh,
		}

		go func() {
			if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Println("err: " + err.Error())
			}
		}()
	}

	<-ctx.Done()
}
