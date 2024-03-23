package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		slog.Error("Missing file name")
		os.Exit(1)
	}

	lines := os.Args[1]

	fmt.Println(lines)
}
