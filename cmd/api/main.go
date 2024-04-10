package main

import (
	"context"
	"math-parser/api/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	r := router.New()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  time.Duration(3) * time.Second,
		WriteTimeout: time.Duration(5) * time.Second,
		IdleTimeout:  time.Duration(5) * time.Second,
	}

	closed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		ctx, cancel :=context.WithTimeout(context.Background(), time.Duration(5) * time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
		}
		close(closed)
	}()
	
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		
	}
	
	<-closed
	
}
