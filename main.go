package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/darkoiv/Kanban-back-go/handlers"
	"github.com/gorilla/mux"
)

func main() {
    // Create server logger
    serverLogger := log.New(os.Stdout, "[Server Log] ", 2)

    // Create router
    router := mux.NewRouter()

    // Create handlers
    lh := handlers.NewListsHandler(serverLogger);

    // Register routes
    getListRoute := router.Methods("GET").Subrouter()
    getListRoute.HandleFunc("/{board}/lists", lh.GetLists)

    // Create server
    server := &http.Server{
        Addr: ":9000",
        Handler: router,
    }

    // Run server
    serverLogger.Println("Starting server!")
    go func() {
        // Check fo server errors
        err := server.ListenAndServe();
        if err != nil {
            serverLogger.Fatalln(err);
        }
    }()

    // Register OS channel, for quit requests
    osChannel := make(chan os.Signal)
    signal.Notify(osChannel, os.Interrupt, os.Kill);

    // Wait for signal
    sig := <-osChannel
    serverLogger.Println("Requestes shutdown, signal:", sig)
}
