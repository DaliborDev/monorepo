package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/DaliborDev/monorepo/pulse/internal"
	"os"
	"time"
)

const defaultTimeTick = 60 * time.Second

type config struct {
	tick time.Duration
}

func main() {
	c := &config{}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		cancel()
	}()

	if err := run(ctx, c); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func (c *config) init(args []string) error {
	/*
		Initialize
		Read flags, config files and environment variables and use them in runtime
	*/
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.String("config", "", "path to .conf file")

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
	err := c.init(os.Args)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.Tick(c.tick):
			{
				//fmt.Printf("Running every %s\n", c.tick)
				internal.ReadLoadAverage()
			}
		}
	}
}
