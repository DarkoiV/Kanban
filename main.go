package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/darkoiv/Kanban-Backend/board"
	"github.com/gorilla/mux"
)

// Main function
func main() {
    serverLogger := log.New(os.Stdout, "[Server Log] ", 2)

    database := connectToDB(serverLogger)
    router  := createRouter(serverLogger, database)

    server := &http.Server{
        Addr: ":9000",
        Handler: router,
        IdleTimeout:  120 * time.Second,
        ReadTimeout:  1   * time.Second,
        WriteTimeout: 1   * time.Second,

    }

    runServer(server, serverLogger);
}

// Connect to DB
func connectToDB(logger *log.Logger) *gorm.DB {
    dsn := "user=kanban_client password=kanban4321 dbname=kanbanDB host=localhost port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        logger.Fatalln("Could not connect to DB")
    }

    logger.Println("Connected to postgres database")

    return db
}

// Creates server router and registers all routes
func createRouter(logger *log.Logger, db *gorm.DB) *mux.Router {
    serverRouter := mux.NewRouter();

    // Register board routes
    bh := board.NewHandler(logger, db);

    boardRoute := serverRouter.PathPrefix("/api/{board}").Subrouter()

    boardRoute.HandleFunc("", bh.GetBoard).Methods("GET")
    boardRoute.HandleFunc("", bh.CreateBoard).Methods("POST")

    boardRoute.HandleFunc("/lists", bh.GetLists).Methods("GET")
    boardRoute.HandleFunc("/lists", bh.UpdateOrder).Methods("PATCH")
    boardRoute.HandleFunc("/list/{list}", bh.DeleteList).Methods("DELETE")

    boardRoute.HandleFunc("/list/{list}", bh.PostCard).Methods("POST")
    boardRoute.HandleFunc("/list/{list}/{task}", bh.UpdateCard).Methods("PUT")
    boardRoute.HandleFunc("/list/{list}/{task}", bh.DeleteCard).Methods("DELETE")

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
