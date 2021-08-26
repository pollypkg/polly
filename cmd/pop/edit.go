package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-clix/cli"
	"github.com/pollypkg/polly/pkg/api/grafana"
	"github.com/pollypkg/polly/pkg/coord"
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

		ctx := context.Background()
		ctx, cancel := coord.WithCancel(ctx)

		p, err := pop.Load(args)
		if err != nil {
			return err
		}

		if err := edit.Check(*p); err != nil {
			return err
		}

		grafanaURL := os.Getenv("GRAFANA_URL")
		if grafanaURL == "" {
			return fmt.Errorf("GRAFANA_URL must be set")
		}
		grafanaToken := os.Getenv("GRAFANA_TOKEN")
		if grafanaToken == "" {
			return fmt.Errorf("GRAFANA_URL must be set")
		}

		c, err := grafana.New(grafanaURL, grafana.Auth{
			Token: grafanaToken,
		})
		if err != nil {
			return err
		}

		srv, err := edit.HTTPHandler(ctx, *p, edit.Opts{Client: c})
		if err != nil {
			return err
		}

		go http.ListenAndServe(":3333", srv)

		<-sigCh
		log.Println("Cleaning up ..")
		cancel()

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
