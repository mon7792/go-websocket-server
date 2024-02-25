package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/mon7792/go-websocket-server/server/srv"
)

func main() {
	log.Println("Starting Server...")

	// Wait for termination signal
	go func() {

		// this will print the number of goroutines.
		for {
			numGoroutines := runtime.NumGoroutine()
			log.Printf("Number of Running Goroutines: %d\n", numGoroutines)
			time.Sleep(1 * time.Second)
		}

	}()

	// context
	_, cancel := context.WithCancel(context.Background())

	// signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	// Start the server
	handler := srv.NewHandler()

	server := srv.NewServer(":8080", handler)
	server.Run()
	<-sig

	// Stop the server
	cancel()
	server.Stop(context.Background())
}
