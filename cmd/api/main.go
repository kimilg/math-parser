package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"math-parser/config"
	"math-parser/internal/math/ports"
	"math-parser/internal/math/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	c := config.New()
	ctx := context.Background()
	app, cleanup := service.NewApplication(ctx, c)
	defer cleanup()
	
	r := ports.New(ports.NewHttpServer(app))
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutWrite,
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
