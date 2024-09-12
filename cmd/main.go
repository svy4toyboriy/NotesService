package main

import (
	"SportsTestWork/config"
	"SportsTestWork/db"
	"SportsTestWork/server"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	cfg := config.Get()

	// Подключение к Postgres
	pgDB, err := db.Dial(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Создание экземпляра сервера
	s := server.Init(ctx, cfg, pgDB)

	// Запуск http-сервера
	addr := ":8080"
	httpServer := &http.Server{
		Addr:         addr,
		Handler:      s,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Running http server on %s\n", addr)

	if err := httpServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
