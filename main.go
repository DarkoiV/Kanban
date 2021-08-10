package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/darkoiv/Kanban-Backend/board"
	"github.com/gorilla/mux"
)


// Main function
func main() {
    serverLogger := log.New(os.Stdout, "[Server Log] ", 2)

    router := createRouter(serverLogger)

    server := &http.Server{
        Addr: ":9000",
        Handler: router,
        IdleTimeout:  120 * time.Second,
        ReadTimeout:  1   * time.Second,
        WriteTimeout: 1   * time.Second,

    }

    runServer(server, serverLogger);
}


// Creates server router and registers all routes
func createRouter(logger *log.Logger) *mux.Router {
    serverRouter := mux.NewRouter();

    // Register board routes
    bh := board.NewHandler(logger);

    boardRoute := serverRouter.PathPrefix("/api/{board}").Subrouter()

    boardRoute.HandleFunc("", bh.GetBoard).Methods("GET")

    boardRoute.HandleFunc("/lists", bh.GetLists).Methods("GET")
    boardRoute.HandleFunc("/lists", bh.UpdateOrder).Methods("PATCH")
    boardRoute.HandleFunc("/list/{id}", bh.DeleteList).Methods("DELETE")

    boardRoute.HandleFunc("/list/{id}", bh.PostCard).Methods("POST")
    boardRoute.HandleFunc("/list/{id}", bh.UpdateCard).Methods("PUT")
    boardRoute.HandleFunc("/list/{id}/{task}", bh.UpdateCard).Methods("DELETE")

    return serverRouter;
}


// Runs server, until requested for quit
func runServer(server *http.Server, serverLogger *log.Logger) {

    // Start server
    serverLogger.Printf("Starting server on port %s . . . ", server.Addr)
    go func() {
        // Check fo server errors
        err := server.ListenAndServe(); if err != nil {
            serverLogger.Println(err);
        }
    }()

    // Register OS channel, for quit requests
    osChannel := make(chan os.Signal)
    signal.Notify(osChannel, os.Interrupt);

    // Wait for signal
    sig := <-osChannel
    fmt.Println();
    serverLogger.Println("Requestes shutdown, signal:", sig)

    // Graceful shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
    if err := server.Shutdown(ctx); err != nil {
        serverLogger.Fatalln("Server didn't shutdown properly:",err)
    }
    defer func(){
        cancel()
    }()

    serverLogger.Println("Server shutdown properly")
}
