package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Define flags
	host := flag.String("host", "127.0.0.1", "The bind address of the server (default: 127.0.0.1)")
	flag.StringVar(host, "h", "127.0.0.1", "The bind address of the server (shorthand for --host)")
	port := flag.Int("port", 3000, "The listening port of the web server (default: 3000)")
	flag.IntVar(port, "p", 3000, "The listening port of the web server (shorthand for --port)")

	// Parse flags
	flag.Parse()

	// Check for directory argument
	if flag.NArg() != 1 {
		fmt.Println("Usage: serve [options] <PATH>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Get the directory path
	dir := flag.Arg(0)

	// Validate directory
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatalf("Error: Directory %s does not exist\n", dir)
	}

	// Define the address
	address := fmt.Sprintf("%s:%d", *host, *port)

	// Serve static files
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	log.Printf("Starting server on %s, serving files from %s", address, dir)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
