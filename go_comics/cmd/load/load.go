package main

import (
	"fmt"
	"io"
	"os"
)

func getOneComic(i int) []byte {
	// fetch metadata about a comic by ID

	return []byte{}
}

func main() {
	var (
		output io.WriteCloser = os.Stdout
		err    error
		// cnt    int
		// fails  int
		// data   []byte
	)

	if len(os.Args) > 1 {
		output, err = os.Create(os.Args[1])

		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading arguments", err)
			os.Exit(1)
		}

		defer output.Close()
	}

	fmt.Println(&output)
}
