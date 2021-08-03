package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/darkoiv/Kanban-back-go/board"
	"github.com/gorilla/mux"
)

func main() {
    // Create server logger
    serverLogger := log.New(os.Stdout, "[Server Log] ", 2)

    // Create router
    router := mux.NewRouter()

    // Create handlers
    bh := board.NewHandler(serverLogger);

    // Register board routes 
    boardRoute := router.PathPrefix("/api/{board}").Subrouter()

    boardRoute.HandleFunc("", bh.GetBoard).Methods("get")
    boardRoute.HandleFunc("/lists", bh.GetLists).Methods("get")

    // Create server
    server := &http.Server{
        Addr: ":9000",
        Handler: router,
        IdleTimeout:  120 * time.Second,
        ReadTimeout:  1   * time.Second,
        WriteTimeout: 1   * time.Second,

    }

    // Run server
    serverLogger.Println("Starting server . . . ")
    go func() {
        // Check fo server errors
        err := server.ListenAndServe(); if err != nil {
            serverLogger.Println(err);
        }
    }()

    // Register OS channel, for quit requests
    osChannel := make(chan os.Signal)
    signal.Notify(osChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT);

    // Wait for signal
    sig := <-osChannel
    serverLogger.Println("Requestes shutdown, signal:", sig)

    // Graceful shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
    if err := server.Shutdown(ctx); err != nil {
        serverLogger.Fatalln("Server didn't shutdown properly:",err)
    }
    defer func(){
        // Some handling here
        cancel()
    }()

    serverLogger.Println("Server shutdown properly")
}
