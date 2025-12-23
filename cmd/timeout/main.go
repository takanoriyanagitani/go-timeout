package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/takanoriyanagitani/go-timeout"
)

func main() {
	var timeoutFlag time.Duration
	flag.DurationVar(&timeoutFlag, "timeout", 10*time.Second, "timeout for the command")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "missing command to execute")
		os.Exit(1)
	}

	builder := timeout.CommandBuilder{
		Name:      args[0],
		Arguments: args[1:],
		Timeout:   timeoutFlag,
	}

	err := builder.Run(context.Background())
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Fprintf(os.Stderr, "command timed out after %v: %s\n", timeoutFlag, args[0])
			os.Exit(124) // Standard exit code for timeout
		}
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
