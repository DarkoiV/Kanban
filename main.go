package main

import (
	"log"
	"net/http"
	"os"

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

    // Start listening
    serverLogger.Println("Starting . . .")
    http.ListenAndServe(":9090", router)
}
