package polly

import "github.com/pollypkg/polly/schema:pollyschema"

pollyschema.PollyPackage

header: {
	name: "test"
	uri:  "github.com/pollypkg/polly/pkg/pop/testdata"
	params: {
        days: 10
    }
}

prometheusAlerts: v0: ExpirySoon: {
	group: "certificates"
	alert: {
		expr: """
              (avg by (exported_namespace, namespace, name) (expiration_timestamp_seconds - time()))
            < (\(header.params.days) * 24 * 3600)
            """
		for:  string | *"1h"
		labels: severity: "warning"
		annotations: {
			summary:     "The cert`{{ $labels.name }}` is {{ $value | humanizeDuration }} from expiry, it should have renewed over a week ago."
			description: "The domain that this cert covers will be unavailable after {{ $value | humanizeDuration }}. Clients using endpoints that this cert protects will start to fail in {{ $value | humanizeDuration }}."
		}
	}
}

prometheusRules: v0: {}

grafanaDashboards: v0: Empty: {
	annotations: list: [
		{
			builtIn:    1
			datasource: "-- Grafana --"
			enable:     true
			hide:       true
			iconColor:  "rgba(0, 211, 255, 1)"
			name:       "Annotations & Alerts"
			target: {
				limit:    100
				matchAny: false
				tags: []
				type: "dashboard"
			}
			type: "dashboard"
		},
	]
	editable:     true
	graphTooltip: 0
	links: []
	panels: []
	schemaVersion: 30
	style:         "dark"
	tags: []
	templating: list: []
	time: {
		from: "now-6h"
		to:   "now"
	}
	timepicker: {}
	title:   "Empty"
	version: 0
}
