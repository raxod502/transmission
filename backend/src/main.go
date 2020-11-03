package main

import (
	"fmt"
	"os"

	"github.com/raxod502/transmission/backend/src/server"
)

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3455"
	}
	addr := host + ":" + port
	fmt.Fprintln(os.Stderr, "transmission: listening on http://"+addr)
	if err := server.Start(addr); err != nil {
		fmt.Fprintf(os.Stderr, "transmission: fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
