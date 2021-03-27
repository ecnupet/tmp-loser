package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ecnu.space/tmp-loser/conf"
	"ecnu.space/tmp-loser/http"
)

const GracefulStopTimeoutSec time.Duration = 5 * time.Second

func main() {
	c := conf.GetAppConfig()
	s := http.InitAndStart(c)
	quitChan := make(chan os.Signal, 1)
	signal.Reset(os.Interrupt, syscall.SIGTERM)
	signal.Notify(quitChan, os.Interrupt, syscall.SIGTERM)
	log.Printf("got signal %s, shutdown server ...", <-quitChan)

	ctx, cancel := context.WithTimeout(context.Background(), GracefulStopTimeoutSec)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown: %v", err)
	}
}
