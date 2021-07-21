package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
    // Create server logger
    serverLogger := log.New(os.Stdout, "[Server Log] ", 2)

    // Start listening
    serverLogger.Println("Starting . . .")
    http.ListenAndServe(":9090", nil)
}
