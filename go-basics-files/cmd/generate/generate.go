package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var (
		filename string
		file     io.WriteCloser = os.Stdout
		err      error
		size     uint
	)

	filename, size = checkArguments(os.Args)

	if !checkIfFileExists(filename) {
		fmt.Println("File doenst exist")
	}

	if file, err = os.Create(filename); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %s\n", err)
		os.Exit(2)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	for i := 0; i < int(size); i++ {
		if _, err := file.Write([]byte("Line number\n")); err != nil {

			os.Exit(2)
		}
	}

	fmt.Println(filename, size)
	fmt.Println(&file)
}

func checkArguments(args []string) (string, uint) {
	var (
		filename string
		size     uint
		sizeTmp  uint64
		err      error
	)
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Missing file name")
		os.Exit(1)
	}

	if len(args) < 3 {
		fmt.Println("Missing generation size")
		os.Exit(1)
	}

	filename = os.Args[1]

	if sizeTmp, err = strconv.ParseUint(os.Args[2], 10, 32); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing unsigned int: %s\n", err)
		os.Exit(3)
	}

	size = uint(sizeTmp)

	return filename, size
}

func checkIfFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}
