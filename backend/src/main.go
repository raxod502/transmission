package main

import (
	"fmt"
	"os"

	"github.com/raxod502/transmission/backend/src/server"
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "localhost:3455"
	}
	fmt.Fprintln(os.Stderr, "transmission: listening on http://"+addr)
	if err := server.Start(addr); err != nil {
		fmt.Fprintf(os.Stderr, "transmission: fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
