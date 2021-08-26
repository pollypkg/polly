package pollyschema

import (
	grafana "github.com/grafana/grafana/cue/data:grafanaschema"
	prometheus "github.com/prometheus/prometheus/schema"
)

// A PollyPackage is a parameterized collection of configuration resources that,
// when provided to a program capable of executing those resources, are
// responsible for collecting, transforming, interpreting, and acting upon a
// particular dataset.  Typically, "a dataset" is understood as the telemetry
// data emitted by a particular system.
//
// @doc(metaschema)
PollyPackage: {
	header: {
		// Simple name of the polly package (node-pop, etc)
		//
		// TODO constrain with regex
		// 
		// @doc(metaschema)
		name: string

		// Fully qualified URI of the polly package (github.com/pollypkg/node-pop)
		//
		// @doc(metaschema)
		uri: string

		// Package-level parameters.
		//
		// @doc(metaschema)
		params: {...}
	}

	// List of signals defined by the polly package.
	//
	// Different signals in the same package may be written in different query
	// languages.
	//
	// NOTE This is a list instead of a struct so that we can allow duplication
	// in the future. That would permit implementations of the "same" signal in
	// different query languages. This would be somewhat analogous to
	// function/method overloading - the consumer's value context can determine
	// whether to use the e.g. promql or flux implementation (assuming both
	// exist).
	//
	// TODO constrain name uniqueness within the list.
	//
	// @doc(metaschema)
	signals?: [Signal, ...Signal]

	// Map of dataface implementations in the polly package.
	//
	// @doc(metaschema)
	datafaces?: [string]: Dataface

	// grafanaDashboards contains definitions of Grafana dashboards that are
	// valid with respect to Grafana dashboard scuemata specifications.
	//
	// @doc(metaschema)
	grafanaDashboards: {
		v0: [string]: (_latest & {arg: grafana.Family.lineages[0]}).out
	}

	// prometheusAlerts contains definitions of Prometheus alerts that are
	// valid with respect to Prometheus alert scuemata specifications.
	//
	// @doc(metaschema)
	prometheusAlerts: {
		v0: [Name=string]: {
			group: string
			alert: (_latest & {arg: prometheus.Alert.lineages[0]}).out & {
				alert: Name
			}
		}
	}

	// prometheusRules contains definitions of Prometheus rules that are
	// valid with respect to Prometheus rule scuemata specifications.
	//
	// @doc(metaschema)
	prometheusRules: {
		v0: [ID=string]: {
			group: string
			rule: (_latest & {arg: prometheus.Rule.lineages[0]}).out & {
				record: ID
			}
		}
	}
}
