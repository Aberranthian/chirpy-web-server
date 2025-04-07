package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// PORT: The port used when serving the site (localhost:PORT).
const PORT = "8080"

func main() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalf("an error occurred getting current working directory: %s", err)
	}

	rootDirectory := filepath.Join(currentDirectory, "..")

	//// Create ServerMux
	mux := http.NewServeMux()

	//// Create Handlers
	fileHandler := http.FileServer(http.Dir(rootDirectory))

	//// Assign Handlers
	mux.Handle("/", fileHandler)

	//// Create and Configure Server
	server := &http.Server{
		Handler: mux,
		Addr:    ":" + PORT,
	}

	fmt.Printf("Starting server in directory \"%s\" on port \"%s\".\n", currentDirectory, PORT)
	log.Fatalf("Server closed with following message: %s", server.ListenAndServe())
}
