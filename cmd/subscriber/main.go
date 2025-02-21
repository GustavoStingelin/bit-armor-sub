package main

import (
	"bit-armor-sub/internal"
	"log"
	"os"
	"os/signal"
)

func main() {
	endpoint := "tcp://127.0.0.1:5555"
	go internal.RunZMQSubscriber(endpoint)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt)
	<-sigterm
	log.Println("Shutting down...")
	os.Exit(0)
}
