package main

import (
	// "SpendingTracker/internal/frontend"
	"SpendingTracker/internal/server"
	"context"
	"os"
	"os/signal"
)

func main() {
	e := server.RunServer()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		e.Logger.Fatal(e.Start(":8080"))
	}()
	<-ctx.Done()

	// go frontend.RunFrontend()
}
