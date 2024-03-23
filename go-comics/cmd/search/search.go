package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// Data sample to be decoded
// "month": "4",
// "num": 571,
// "link": "",
// "year": "2009",
// "news": "",
// "safe_title": "Can't Sleep",
// "transcript": "Lorem ipsum dolor sit amet",
// "img": "https://imgs.xkcd.com/comics/cant_sleep.png",
// "title": "Can't Sleep",
// "day": "20"

type xkcd struct {
	Num        int    `json:num`
	Day        string `json:day`
	Month      string `json:month`
	Year       string `json:year`
	Title      string `json:title`
	Transcript string `json:transcript`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "no file given")
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "no search term given")
		os.Exit(0)
	}

	file := os.Args[1]

	var (
		items []xkcd
		terms []string
		input io.ReadCloser
		count int
		err   error
	)

	if input, err = os.Open(file); err != nil {
		fmt.Fprintf(os.Stderr, "could not open file: %s\n", err)
		os.Exit(1)
	}

	// decode file
	if err = json.NewDecoder(input).Decode(&items); err != nil {
		fmt.Fprintln(os.Stderr, "invalid json: %s\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "read %d comics\n", len(items))

	// get search terms
	for _, t := range os.Args[2:] {
		terms = append(terms, strings.ToLower(t))
	}

	// search!
outer:
	for _, item := range items {
		title := strings.ToLower(item.Title)
		transcript := strings.ToLower(item.Transcript)

		for _, term := range terms {
			if !strings.Contains(title, term) && !strings.Contains(transcript, term) {
				continue outer
			}
		}

		fmt.Fprintf(os.Stdout, "https://xkcd.com/%d %s/%s/%s %q\n", item.Num, item.Day, item.Month, item.Year, item.Title)
		count++
	}

	fmt.Fprintln(os.Stdout, "Success! Exiting...")
}
