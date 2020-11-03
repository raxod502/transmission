package main

import (
	"fmt"
	"os"

	"github.com/raxod502/transmission/backend/server"
)

func main() {
	fmt.Fprintln(os.Stderr, "transmission: listening on http://localhost:3455")
	if err := server.Start("localhost:3455"); err != nil {
		fmt.Fprintf(os.Stderr, "transmission: fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
