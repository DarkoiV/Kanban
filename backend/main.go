package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

    "github.com/joho/godotenv"

	"github.com/darkoiv/Kanban/backend/board"
	"github.com/darkoiv/Kanban/backend/front"
	"github.com/gorilla/mux"
)

// Main function
func main() {
    logger := log.New(os.Stdout, "[Server Log] ", 2)

    loadENV(logger)
    database := connectToDB(logger)
    router  := createRouter(logger, database)

    server := &http.Server{
        Addr: ":9000",
        Handler: router,
        IdleTimeout:  120 * time.Second,
        ReadTimeout:  1   * time.Second,
        WriteTimeout: 1   * time.Second,

    }

    runServer(server, logger);
}

// Load ENV
func loadENV (logger *log.Logger) {
    if err := godotenv.Load(); err != nil {
        logger.Fatalln("Error loading .env", err)
    }
}

// Connect to DB
func connectToDB(logger *log.Logger) *gorm.DB {
    var db *gorm.DB
    var err error

    dbDialect := os.Getenv("DIALECT")

    switch dbDialect {
        case "sqlite":
            db, err = gorm.Open(sqlite.Open("local.db"), &gorm.Config{})

        case "postgres":
            dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
                    os.Getenv("DB_USER"),
                    os.Getenv("DB_PASS"),
                    os.Getenv("DB_NAME"),
                    os.Getenv("DB_HOST"),
                    os.Getenv("DB_PORT"))
            db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

        default:
            logger.Fatalln("Unknown database dialect:", dbDialect)
    }


    if err != nil || db == nil {
        logger.Fatalln("Could not connect to DB", err)
    }

    logger.Println("Connected to", dbDialect, "database")

    return db
}

// Creates server router and registers routes
func createRouter(logger *log.Logger, db *gorm.DB) *mux.Router {
    serverRouter := mux.NewRouter();

    // Create API route
    apiRoute := serverRouter.PathPrefix("/api").Subrouter()

    // Create board handler
    bh := board.NewHandler(logger, db);
    apiRoute.HandleFunc("/new", bh.CreateBoard).Methods("POST")

    // Create board specific route
    boardRoute := apiRoute.PathPrefix("/{boardID:[0-9]+}").Subrouter()
    boardRoute.HandleFunc("", bh.GetBoard).Methods("GET")
    boardRoute.HandleFunc("", bh.DeleteBoard).Methods("DELETE")
    boardRoute.HandleFunc("/lists", bh.UpdateLists).Methods("PATCH")

    // Create list specific route
    listRoute := boardRoute.PathPrefix("/{listID:[0-9]+}").Subrouter()
    listRoute.HandleFunc("", bh.DeleteList).Methods("DELETE")
    listRoute.HandleFunc("", bh.PostTask).Methods("POST")

    listRoute.HandleFunc("/{taskID:[0-9]+}", bh.UpdateTask).Methods("PUT")
    listRoute.HandleFunc("/{taskID:[0-9]+}", bh.DeleteTask).Methods("DELETE")

    // Catch bad API calls
    serverRouter.PathPrefix("/api").HandlerFunc(func(rw http.ResponseWriter, rq *http.Request){
        http.Error(rw, "", http.StatusBadRequest)
    })

    // Serve single page app
    fh := front.NewHandler(logger, "dist", "index.html")
    serverRouter.PathPrefix("/").Handler(fh).Methods("GET")

    return serverRouter;
}

// Runs server, until requested for quit
func runServer(server *http.Server, logger *log.Logger) {

    // Start server
    logger.Printf("Starting server on port %s . . . ", server.Addr)
    go func() {
        // Check fo server errors
        err := server.ListenAndServe(); if err != nil {
            logger.Println(err);
        }
    }()

    // Register OS channel, for quit requests
    osChannel := make(chan os.Signal)
    signal.Notify(osChannel, os.Interrupt);

    // Wait for signal
    sig := <-osChannel
    fmt.Println();
    logger.Println("Requestes shutdown, signal:", sig)

    // Graceful shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
    if err := server.Shutdown(ctx); err != nil {
        logger.Fatalln("Server didn't shutdown properly:",err)
    }
    defer func(){
        cancel()
    }()

    logger.Println("Server shutdown properly")
}
