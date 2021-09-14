package main

import (
	"log"
	"os"

	"cuelang.org/go/cue/errors"
	"github.com/go-clix/cli"
)

func main() {
	log.SetFlags(0)

	cmd := cli.Command{
		Use: "pop",
	}

	cmd.AddCommand(
		mixCmd(),
	)

	if err := cmd.Execute(); err != nil {
		errors.Print(os.Stderr, err, nil)
		os.Exit(1)
	}
}
