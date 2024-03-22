package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func getOneComic(i int) []byte {
	// fetch metadata about a comic by ID

	return nil
}

func main() {
	var (
		output  io.WriteCloser = os.Stdout
		err     error
		counter int
		fails   int
		data    []byte
	)

	if len(os.Args) > 1 {
		output, err = os.Create(os.Args[1])

		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading arguments", err)
			os.Exit(1)
		}

		defer output.Close()
	}

	fmt.Println("[")
	defer fmt.Println("]")

	for i := 1; fails < 2; i++ {
		if data = getOneComic(i); data == nil {
			// if there is no data, increment fails
			fails++
			continue
		}

		if counter > 0 {
			fmt.Fprint(output, ",")
		}

		_, err := io.Copy(output, bytes.NewBuffer(data))

		if err != nil {
			fmt.Fprintf(os.Stderr, "stopped: %s\n", err)
			os.Exit(1)
		}

		fails = 0
		counter++
	}

	fmt.Fprintf(os.Stdout, "read %d comics\n", counter)
}
