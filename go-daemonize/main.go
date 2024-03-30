package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"
)

const defaultTimeTick = 60 * time.Second

type config struct {
	tick time.Duration
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	c := &config{}

	defer func() {
		cancel()
	}()

	if err := run(ctx, c); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func (c *config) init(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	var (
		tick = flags.Duration("interval", defaultTimeTick, "Time tick")
	)

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "Greska pri parsiranju", err)
		return err
	}

	c.tick = *tick

	return nil
}

func run(ctx context.Context, c *config) error {
	c.init(os.Args)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.Tick(c.tick):
			fmt.Printf("Running every %s\n", c.tick)
		}
	}
}
