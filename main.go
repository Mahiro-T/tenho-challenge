package main

import (
	"flag"
	"fmt"
	"go.kyoto.codes/v3/rendering"
	"net/http"
)

func setup() *http.ServeMux {
	// Initialize mux
	mux := http.NewServeMux()

	// Setup assets
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Setup pages
	mux.Handle("/", rendering.Handler(IndexPage))

	// Setup components (htmx)

	// Return mux
	return mux
}

func main() {
	// Parse arguments
	addr := flag.String("http", ":8000", "Serving address")
	flag.Parse()

	// Serve
	fmt.Printf("Serving on %s\n", *addr)
	if err := http.ListenAndServe(*addr, setup()); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
