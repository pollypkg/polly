package main

import (
	"github.com/go-clix/cli"
	"github.com/pollypkg/polly/pkg/pop"
)

func mixCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "mix <path>",
		Short: "output pop in mixin compatible format",
		Args:  cli.ArgsMin(1),
	}

	printer := cmd.Flags().StringP("output", "o", "json", "output format. One of json, yaml")
	system := cmd.Flags().StringP("system", "s", "", "choose subsystem. One of alerts, rules, grafana (default all)")

	cmd.Run = func(cmd *cli.Command, args []string) error {
		p, err := pop.Load(args)
		if err != nil {
			return err
		}

		mix, err := p.Mixin()
		if err != nil {
			return err
		}

		var out interface{} = mix
		switch *system {
		case "alerts":
			out = map[string]interface{}{"groups": mix.PrometheusAlerts}
		case "rules":
			out = map[string]interface{}{"groups": mix.PrometheusRules}
		case "grafana":
			out = mix.GrafanaDashboards
		}

		return choosePrinter(*printer).Print(out)
	}

	return cmd
}
