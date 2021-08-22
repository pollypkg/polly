package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"os/signal"

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
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt)

		p, err := pop.Load(args)
		if err != nil {
			return err
		}

		if err := edit.Check(*p); err != nil {
			return err
		}

		grafanaURL := "http://localhost:3000"
		grafanaToken := "eyJrIjoiRTJzVzZIS2RsNlhvSTJVNXlEWUM3RDNwa1JNanRQNjkiLCJuIjoidGVzdCIsImlkIjoxfQ=="
		c, err := grafana.New(grafanaURL, grafana.Auth{
			Token: grafanaToken,
		})
		if err != nil {
			return err
		}

		e, err := edit.Edit(*p, edit.Opts{Client: c})
		if err != nil {
			return err
		}

		dbs := p.Dashboards()
		if len(dbs) != 1 {
			var strs []string
			for _, d := range dbs {
				strs = append(strs, d.Name())
			}

			return fmt.Errorf("Editing alpha requires exactly one dashboard, found %s", strs)
		}

		if err := e.Add(dbs[0].Name()); err != nil {
			return err
		}

		<-sigCh
		log.Println("Cleaning up ..")
		if err := e.Close(); err != nil {
			return err
		}

		return nil
	}

	return cmd
}

// pop-<hostname>-<name>
func dashboardID(name string) string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	if len(hostname) > 11 {
		hostname = hostname[:11]
	}

	hash := base64.RawStdEncoding.EncodeToString(sha256.New().Sum([]byte(name)))[:28]
	id := fmt.Sprintf("pop-%s-%s", hostname, hash)
	if len(id) != 40 {
		panic(len(id))
	}

	return id
}
