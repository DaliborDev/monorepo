package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type ByteSliceCallback func(int) []byte

func getOneComic(i int) []byte {
	// fetch metadata about a comic by ID
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "cant read from url: %s\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "skipping %d, with code %d\n", i, resp.StatusCode)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "something went wrong reading the body: %s\n", err)
	}

	return body
}

func main() {
	startTime := makeTimestamp()
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
	}

	fmt.Println("[")
	defer fmt.Println("]")

	for i := 1; fails < 2; i++ {
		data = benchmarkingWrapper(getOneComic, i)
		if data == nil {
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

	output.Close()

	endTime := makeTimestamp()

	fmt.Printf("Total time of execution: %d\n", endTime-startTime)
}

func benchmarkingWrapper(callback ByteSliceCallback, i int) []byte {
	startTime := makeTimestamp()
	data := callback(i)
	endTime := makeTimestamp()
	fmt.Fprintf(os.Stdout, "request %d took %dms\n", i, endTime-startTime)

	return data
}

func makeTimestamp() int64 {
	// https://stackoverflow.com/a/34860368
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}
