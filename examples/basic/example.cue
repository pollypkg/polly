package pollyexample

import "github.com/pollypkg/polly/schema:pollyschema"

// NOTE we could make it an expectation that the emit value is
// a pollyschema.PollyPackage
examplepkg: pollyschema.PollyPackage & {
	header: {
		name: "node-exporter"
		uri:  "github.com/pollypkg/polly/examples/basic"
		params: {}
	}
	signals: [
		{
			"NumCpu": {
				params: {job: string, instance: string}
				lang: "promql"
				expr: "count without (cpu) (count without (mode) (node_cpu_seconds_total{job=\"\(args.job)\", instance=\"\(args.instance)\"}))"
			}
		},
	]

	// NOTE examples here are gonna be pretty broken until we have
	// a reliable Grafana schema
	grafanaDashboards: v0: nodedashboard: {
		uid: "a8b327a" // Define the uid that Grafana will internally use to uniquely identify the dashboard

		// ...in a real dashboard definition, there will actually be stuff here! Though, probably not too much;
		// that's the value of relying on schema-defined defaults.
	}
}
