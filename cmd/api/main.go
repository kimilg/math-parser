package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"math-parser/api/resource/math"
	"math-parser/api/router"
	"math-parser/config"
	"math-parser/db"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const fmtDBUrl = "postgres://%s:%s@%s:%d/%s?sslmode=%s"
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	c := config.New()
	ctx := context.Background()

	dbString := fmt.Sprintf(fmtDBUrl, c.DB.Username, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.DBName, false)
	conn, err := pgx.Connect(ctx, dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	
	queries := db.New(conn)
	repository := math.NewRepository(queries)
	r := router.New(repository)

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
