package prometheusOperator

import "github.com/pollypkg/polly/schema:pollyschema"

// NOTE we could make it an expectation that the emit value is
// a pollyschema.PollyPackage
prometheusOperator: pollyschema.PollyPackage & {
	header: {
		name: "prometheus-operator"
		uri:  "github.com/pollypkg/polly/examples/prometheus-operator"
		params: {}
	}
	signals: [
		// List Errors
		{
			name: "ListErrors"
			lang: "promql"
			params: {job: string}
			query: """
                sum by (controller,namespace) (rate(prometheus_operator_list_operations_failed_total{job="\(params.job)"}[10m]))
                /
                sum by (controller,namespace) (rate(prometheus_operator_list_operations_total{job="\(params.job)"}[10m]))
                > 0.4
            """
		},
		// WatchErrors
		{
			name: "WatchErrors"
			lang: "promql"
			params: {job: string}
			query: """
              sum by (controller,namespace) (rate(prometheus_operator_watch_operations_failed_total{job="\(params.job)"}[10m]))
              /
              sum by (controller,namespace) (rate(prometheus_operator_watch_operations_total{job="\(params.job)"}[10m]))
              > 0.4
            """
		},
		// SyncFailed
		{
			name: "SyncFailed"
			lang: "promql"
			params: {job: string}
			query: """
              min_over_time(prometheus_operator_syncs{status=\"failed\",job="\(params.job)"}[5m]) > 0
            """
		},
		// Reconcile Errors
		{
			name: "ReconcileErrors"
			lang: "promql"
			params: {job: string}
			query: """
                (sum by (controller,namespace) (rate(prometheus_operator_reconcile_errors_total{job="\(params.job)"}[5m])))
                /
                (sum by (controller,namespace) (rate(prometheus_operator_reconcile_operations_total{job="\(params.job)"}[5m]))) 
                > 0.1
            """
		},
		// NodeLookupErrors
		{
			name: "NodeLookupErrors"
			lang: "promql"
			params: {job: string}
			query: """
				rate(prometheus_operator_node_address_lookup_errors_total{job="\(params.job)"}[5m]) > 0.1
			"""
		},
		// NotReady
		{
			name: "NotReady"
			lang: "promql"
			params: {job: string}
			query: """
				min by(namespace, controller) (max_over_time(prometheus_operator_ready{job="\(params.job)"}[5m]) == 0)
			"""
		},
		// RejectedResources
		{
			name: "RejectedResources"
			lang: "promql"
			params: {job: string}
			query: """
				min_over_time(prometheus_operator_managed_resources{state=\"rejected\",job="\(params.job)"}[5m]) > 0
			"""
		},
	]

	datafaces: {
	}

	prometheusAlerts: v0: prometheusOperator: {

	}

	grafanaDashboards: v0: nodedashboard: {
		uid: "a8b327a" // Define the uid that Grafana will internally use to uniquely identify the dashboard
	}
}
