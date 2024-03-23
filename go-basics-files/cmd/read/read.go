package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		file []byte
		err  error
	)

	if file, err = os.ReadFile(os.Args[1]); err != nil {
		log.Fatalf("unable to read file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(file))
}
