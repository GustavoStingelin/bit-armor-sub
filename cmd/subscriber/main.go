package main

import (
	"bit-armor-sub/internal"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}
	endpoint := "tcp://127.0.0.1:5555"

	internal.InitDBPool()
	go internal.RunZMQSubscriber(endpoint)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt)
	<-sigterm
	slog.Info("Shutting down...")
	os.Exit(0)
}
