package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-clix/cli"
	"github.com/pollypkg/polly/pkg/api/grafana"
	"github.com/pollypkg/polly/pkg/edit"
	"github.com/pollypkg/polly/pkg/pop"
)

func editCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "edit <path> [resource]",
		Short: "interactive editing session",
		Args:  cli.ArgsMin(1),
	}

	cmd.Run = func(cmd *cli.Command, args []string) error {
		p, err := pop.Load(args)
		if err != nil {
			return err
		}

		if err := edit.Check(*p); err != nil {
			return err
		}

		return nil

		c, err := grafana.New("http://localhost:3000", grafana.Auth{
			Token: "eyJrIjoieEpyUTl4SUQ4ZVBGaWlGT0RleHhvYlZrRmxLZmo4d24iLCJuIjoidGVzdCIsImlkIjoxfQ==",
		})
		if err != nil {
			return err
		}

		w, err := c.Watcher()
		if err != nil {
			return err
		}

		if err := w.Add("P5bV68M7z", func(i interface{}) error {
			data, err := json.MarshalIndent(i, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(data))
			return nil
		}); err != nil {
			return err
		}

		select {}
	}

	return cmd
}
