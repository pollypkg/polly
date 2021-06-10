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
		// Number of CPUs the node has
		{
			name: "NumCpu"
			lang: "promql"
			params: {job: string, instance: string}
			query: "count without (cpu) (count without (mode) (node_cpu_seconds_total{job=\"\(params.job)\", instance=\"\(params.instance)\"}))"
		},
		// Amount of memory currently in use
		{
			name: "MemoryUtilization"
			lang: "promql"
			params: {job: string, instance: string}
			query: "1 - (node_memory_MemAvailable_bytes{job=\"\(params.job)\", instance=\"\(params.instance)\"} / node_memory_MemTotal_bytes{job=\"\(params.job)\", instance=\"\(params.instance)\"})"
		},
		// One minute rate of major page faults
		{
			name: "VmstatPGMajFault"
			lang: "promql"
			params: {job: string, instance: string}
			query: "rate(node_vmstat_pgmajfault{job=\"\(params.job)\", instance=\"\(params.instance)\"}[1m])"
		},
	]

	datafaces: {
		"use_mem": pollyschema.USE & {
			frames: {
				utilization: "MemoryUtilization"
				saturation:  "VmstatPGMajFault"
				errors:      "" // NOTE omitting/empty string for this would be an error
			}
		}
	}

	// NOTE examples here are gonna be pretty broken/artificial until we have a
	// reliable Grafana schema
	grafanaDashboards: v0: nodedashboard: {
		uid: "a8b327a" // Define the uid that Grafana will internally use to uniquely identify the dashboard

		// ...in a real dashboard definition, there will actually be stuff here! Though, probably not too much;
		// that's the value of relying on schema-defined defaults.
	}
}
