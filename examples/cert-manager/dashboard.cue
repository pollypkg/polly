package polly

grafanaDashboards: v0: {
	// TODO Camel casing of these names? Maybe yet another thing to establish practices around
	CertManager: {
		annotations: {
			list: [{
				builtIn:    1
				datasource: "-- Grafana --"
				enable:     true
				hide:       true
				iconColor:  "rgba(0, 211, 255, 1)"
				name:       "Annotations & Alerts"
				type:       "dashboard"
			}]
		}
		description:  ""
		editable:     true
		graphTooltip: 1
		id:           59
		iteration:    1616445892702
		links: []
		panels: [{
			datasource:  "$datasource"
			description: "The number of certificates in the ready state."
			fieldConfig: {
				defaults: {
					custom: {}
					mappings: []
					thresholds: {
						mode: "absolute"
						steps: [{
							color: "green"
						}, {
							color: "red"
							value: 1
						}]
					}
				}
				overrides: [{
					matcher: {
						id:      "byName"
						options: "True"
					}
					properties: [{
						id: "thresholds"
						value: {
							mode: "absolute"
							steps: [{
								color: "green"
							}]
						}
					}]
				}]
			}
			gridPos: {
				h: 8
				w: 12
				x: 0
				y: 0
			}
			id: 2
			options: {
				colorMode:   "value"
				graphMode:   "area"
				justifyMode: "auto"
				orientation: "auto"
				reduceOptions: {
					calcs: ["lastNotNull"]
					fields: ""
					values: false
				}
				text: {}
				textMode: "auto"
			}
			pluginVersion: "7.4.5"
			targets: [{
				expr:         "sum by (condition) (certmanager_certificate_ready_status)"
				interval:     ""
				legendFormat: "{{condition}}"
				refId:        "A"
			}]
			title: "Certificates Ready"
			type:  "stat"
		}, {
			datasource: "$datasource"
			fieldConfig: {
				defaults: {
					custom: {}
					decimals: 1
					mappings: []
					thresholds: {
						mode: "absolute"
						steps: [{
							color: "red"
						}, {
							color: "#EAB839"
							value: 604800
						}, {
							color: "green"
							value: 1209600
						}]
					}
					unit: "dtdurations"
				}
				overrides: []
			}
			gridPos: {
				h: 8
				w: 12
				x: 12
				y: 0
			}
			id: 4
			options: {
				colorMode:   "value"
				graphMode:   "area"
				justifyMode: "auto"
				orientation: "auto"
				reduceOptions: {
					calcs: ["lastNotNull"]
					fields: ""
					values: false
				}
				text: {}
				textMode: "auto"
			}
			pluginVersion: "7.4.5"
			targets: [{
				expr:         "min(certmanager_certificate_expiration_timestamp_seconds) - time()"
				hide:         false
				instant:      true
				interval:     ""
				legendFormat: ""
				refId:        "A"
			}, {
				expr:         "vector(1250000)"
				hide:         true
				instant:      true
				interval:     ""
				legendFormat: ""
				refId:        "B"
			}]
			title: "Soonest Cert Expiry"
			type:  "stat"
		}, {
			datasource:  "$datasource"
			description: "Status of the certificates. Values are True, False or Unknown."
			fieldConfig: {
				defaults: {
					custom: {
						filterable: false
					}
					mappings: [{
						from:     ""
						id:       0
						operator: ""
						text:     "Yes"
						to:       ""
						type:     1
						value:    ""
					}]
					thresholds: {
						mode: "absolute"
						steps: [{
							color: "green"
						}, {
							color: "red"
							value: 80
						}]
					}
					unit: "none"
				}
				overrides: [{
					matcher: {
						id:      "byName"
						options: "Ready Status"
					}
					properties: [{
						id:    "custom.width"
						value: 148
					}]
				}, {
					matcher: {
						id:      "byName"
						options: "Valid Until"
					}
					properties: [{
						id:    "custom.width"
						value: 203
					}]
				}, {
					matcher: {
						id:      "byName"
						options: "Namespace"
					}
					properties: [{
						id:    "custom.width"
						value: 324
					}]
				}, {
					matcher: {
						id:      "byName"
						options: "Valid Until"
					}
					properties: [{
						id:    "unit"
						value: "dateTimeAsIso"
					}]
				}]
			}
			gridPos: {
				h: 8
				w: 12
				x: 0
				y: 8
			}
			id: 9
			options: {
				showHeader: true
				sortBy: [{
					desc:        false
					displayName: "Valid Until"
				}]
			}
			pluginVersion: "7.4.5"
			targets: [{
				expr:         "avg by (name, condition) (certmanager_certificate_ready_status == 1)"
				format:       "table"
				instant:      true
				interval:     ""
				legendFormat: ""
				refId:        "A"
			}, {
				expr:         "avg by (name, namespace, exported_namespace) (certmanager_certificate_expiration_timestamp_seconds) * 1000"
				format:       "table"
				instant:      true
				interval:     ""
				legendFormat: ""
				refId:        "B"
			}]
			title: "Certificates"
			transformations: [{
				id: "seriesToColumns"
				options: {
					byField: "name"
				}
			}, {
				id: "organize"
				options: {
					excludeByName: {
						Time:               true
						"Time 1":           true
						"Time 2":           true
						"Value #A":         true
						exported_namespace: false
					}
					indexByName: {
						Time:               6
						"Value #A":         5
						"Value #B":         4
						condition:          3
						exported_namespace: 1
						name:               2
						namespace:          0
					}
					renameByName: {
						"Value #B":         "Valid Until"
						condition:          "Ready Status"
						exported_namespace: "Certificate Namespace"
						name:               "Certificate"
						namespace:          "Namespace"
					}
				}
			}]
			type: "table"
		}, {
			aliasColors: {}
			bars:        false
			dashLength:  10
			dashes:      false
			datasource:  "$datasource"
			description: "The rate of controller sync requests."
			fieldConfig: {
				defaults: {
					custom: {}
					links: []
				}
				overrides: []
			}
			fill:         1
			fillGradient: 0
			gridPos: {
				h: 8
				w: 12
				x: 12
				y: 8
			}
			hiddenSeries: false
			id:           7
			interval:     "20s"
			legend: {
				avg:     false
				current: false
				max:     false
				min:     false
				show:    true
				total:   false
				values:  false
			}
			lines:         true
			linewidth:     1
			maxDataPoints: 250
			nullPointMode: "null"
			options: {
				alertThreshold: true
			}
			percentage:    false
			pluginVersion: "7.4.5"
			pointradius:   2
			points:        false
			renderer:      "flot"
			seriesOverrides: []
			spaceLength: 10
			stack:       false
			steppedLine: false
			targets: [{
				expr: """
					sum by (controller) (
					rate(certmanager_controller_sync_call_count[$__rate_interval])
					)
					"""
				interval:     ""
				legendFormat: "{{controller}}"
				refId:        "A"
			}]
			thresholds: []
			timeRegions: []
			title: "Controller Sync Requests/sec"
			tooltip: {
				shared:     true
				sort:       0
				value_type: "individual"
			}
			type: "graph"
			xaxis: {
				mode: "time"
				show: true
				values: []
			}
			yaxes: [{
				format:  "reqps"
				logBase: 1
				min:     "0"
				show:    true
			}, {
				format:  "short"
				logBase: 1
				show:    true
			}]
			yaxis: {
				align: false
			}
		}, {
			aliasColors: {}
			bars:        false
			dashLength:  10
			dashes:      false
			datasource:  "$datasource"
			description: "Rate of requests to ACME provider."
			fieldConfig: {
				defaults: {
					custom: {}
					links: []
				}
				overrides: []
			}
			fill:         1
			fillGradient: 0
			gridPos: {
				h: 8
				w: 12
				x: 0
				y: 16
			}
			hiddenSeries: false
			id:           6
			interval:     "20s"
			legend: {
				avg:       false
				current:   false
				hideEmpty: true
				hideZero:  false
				max:       false
				min:       false
				show:      true
				total:     false
				values:    false
			}
			lines:         true
			linewidth:     1
			maxDataPoints: 250
			nullPointMode: "null"
			options: {
				alertThreshold: true
			}
			percentage:    false
			pluginVersion: "7.4.5"
			pointradius:   2
			points:        false
			renderer:      "flot"
			seriesOverrides: []
			spaceLength: 10
			stack:       false
			steppedLine: false
			targets: [{
				expr: """
					sum by (method, path, status) (
					rate(certmanager_http_acme_client_request_count[$__rate_interval])
					)
					"""
				interval:     ""
				legendFormat: "{{method}} {{path}} {{status}}"
				refId:        "A"
			}]
			thresholds: []
			timeFrom: null
			timeRegions: []
			timeShift: null
			title:     "ACME HTTP Requests/sec"
			tooltip: {
				shared:     true
				sort:       0
				value_type: "individual"
			}
			type: "graph"
			xaxis: {
				buckets: null
				mode:    "time"
				name:    null
				show:    true
				values: []
			}
			yaxes: [{
				format:  "reqps"
				label:   null
				logBase: 1
				max:     null
				min:     "0"
				show:    true
			}, {
				format:  "short"
				label:   null
				logBase: 1
				max:     null
				min:     null
				show:    true
			}]
			yaxis: {
				align:      false
				alignLevel: null
			}
		}, {
			aliasColors: {}
			bars:        false
			dashLength:  10
			dashes:      false
			datasource:  "$datasource"
			description: "Average duration of requests to ACME provider. "
			fieldConfig: {
				defaults: {
					custom: {}
					links: []
				}
				overrides: []
			}
			fill:         1
			fillGradient: 0
			gridPos: {
				h: 8
				w: 12
				x: 12
				y: 16
			}
			hiddenSeries: false
			id:           10
			interval:     "30s"
			legend: {
				avg:       false
				current:   false
				hideEmpty: true
				hideZero:  false
				max:       false
				min:       false
				show:      true
				total:     false
				values:    false
			}
			lines:         true
			linewidth:     1
			maxDataPoints: 250
			nullPointMode: "null"
			options: {
				alertThreshold: true
			}
			percentage:    false
			pluginVersion: "7.4.5"
			pointradius:   2
			points:        false
			renderer:      "flot"
			seriesOverrides: []
			spaceLength: 10
			stack:       false
			steppedLine: false
			targets: [{
				expr: """
					sum by (method, path, status) (rate(certmanager_http_acme_client_request_duration_seconds_sum[$__rate_interval]))
					/
					sum by (method, path, status) (rate(certmanager_http_acme_client_request_duration_seconds_count[$__rate_interval]))
					"""
				interval:     ""
				legendFormat: "{{method}} {{path}} {{status}}"
				refId:        "A"
			}]
			thresholds: []
			timeFrom: null
			timeRegions: []
			timeShift: null
			title:     "ACME HTTP Request avg duration"
			tooltip: {
				shared:     true
				sort:       0
				value_type: "individual"
			}
			type: "graph"
			xaxis: {
				buckets: null
				mode:    "time"
				name:    null
				show:    true
				values: []
			}
			yaxes: [{
				format:  "s"
				label:   null
				logBase: 1
				max:     null
				min:     "0"
				show:    true
			}, {
				format:  "short"
				label:   null
				logBase: 1
				max:     null
				min:     null
				show:    true
			}]
			yaxis: {
				align:      false
				alignLevel: null
			}
		}, {
			aliasColors: {
				max: "dark-yellow"
			}
			bars:        false
			dashLength:  10
			dashes:      false
			datasource:  "$datasource"
			description: "CPU Usage and limits, as percent of a vCPU core. "
			fieldConfig: {
				defaults: {
					custom: {}
					links: []
				}
				overrides: []
			}
			fill:         0
			fillGradient: 0
			gridPos: {
				h: 8
				w: 6
				x: 0
				y: 24
			}
			hiddenSeries: false
			id:           12
			interval:     "1m"
			legend: {
				avg:     false
				current: false
				max:     false
				min:     false
				show:    true
				total:   false
				values:  false
			}
			lines:     true
			linewidth: 1
			links: []
			maxDataPoints: 250
			nullPointMode: "null"
			options: {
				alertThreshold: true
			}
			percentage:    false
			pluginVersion: "7.4.5"
			pointradius:   2
			points:        false
			renderer:      "flot"
			seriesOverrides: [{
				alias:        "CPU"
				fill:         1
				fillGradient: 5
			}, {
				alias:  "/Request.*/"
				color:  "#FF9830"
				dashes: true
			}, {
				alias:  "/Limit.*/"
				color:  "#F2495C"
				dashes: true
			}]
			spaceLength: 10
			stack:       false
			steppedLine: false
			targets: [{
				expr:           "avg by (pod) (rate(container_cpu_usage_seconds_total{container=\"cert-manager\"}[$__rate_interval]))"
				format:         "time_series"
				hide:           false
				interval:       ""
				intervalFactor: 2
				legendFormat:   "CPU {{pod}}"
				refId:          "A"
			}, {
				expr:           "avg by (pod) (kube_pod_container_resource_limits_cpu_cores{container=\"cert-manager\"})"
				format:         "time_series"
				hide:           true
				interval:       ""
				intervalFactor: 1
				legendFormat:   "Limit {{pod}}"
				refId:          "B"
			}, {
				expr:           "avg by (pod) (kube_pod_container_resource_requests_cpu_cores{container=\"cert-manager\"})"
				format:         "time_series"
				hide:           true
				interval:       ""
				intervalFactor: 1
				legendFormat:   "Request {{pod}}"
				refId:          "C"
			}]
			thresholds: []
			timeFrom: null
			timeRegions: []
			timeShift: null
			title:     "CPU"
			tooltip: {
				shared:     true
				sort:       0
				value_type: "individual"
			}
			type: "graph"
			xaxis: {
				buckets: null
				mode:    "time"
				name:    null
				show:    true
				values: []
			}
			yaxes: [{
				format:  "percentunit"
				label:   null
				logBase: 1
				max:     null
				min:     "0"
				show:    true
			}, {
				format:  "short"
				label:   null
				logBase: 1
				max:     null
				min:     null
				show:    true
			}]
			yaxis: {
				align:      false
				alignLevel: null
			}
		}, {
			aliasColors: {
				max: "dark-yellow"
			}
			bars:        false
			dashLength:  10
			dashes:      false
			datasource:  "$datasource"
			description: "Percent of the time that the CPU is being throttled. Higher is badderer. "
			fieldConfig: {
				defaults: {
					custom: {}
					links: []
				}
				overrides: []
			}
			fill:         0
			fillGradient: 0
			gridPos: {
				h: 8
				w: 6
				x: 6
				y: 24
			}
			hiddenSeries: false
			id:           14
			interval:     "1m"
			legend: {
				avg:     false
				current: false
				max:     false
				min:     false
				show:    true
				total:   false
				values:  false
			}
			lines:     true
			linewidth: 1
			links: []
			maxDataPoints: 250
			nullPointMode: "connected"
			options: {
				alertThreshold: true
			}
			percentage:    false
			pluginVersion: "7.4.5"
			pointradius:   2
			points:        false
			renderer:      "flot"
			seriesOverrides: [{
				alias:        "/external-dns.*/"
				fill:         1
				fillGradient: 5
			}]
			spaceLength: 10
			stack:       false
			steppedLine: false
			targets: [{
				expr: """
					avg by (pod) (
					rate(container_cpu_cfs_throttled_periods_total{container="cert-manager"}[$__rate_interval])
					/
					rate(container_cpu_cfs_periods_total{container="cert-manager"}[$__rate_interval])
					)
					"""
				format:         "time_series"
				hide:           false
				interval:       ""
				intervalFactor: 2
				legendFormat:   "{{pod}}"
				refId:          "A"
			}]
			thresholds: []
			timeFrom: null
			timeRegions: []
			timeShift: null
			title:     "CPU Throttling"
			tooltip: {
				shared:     true
				sort:       0
				value_type: "individual"
			}
			type: "graph"
			xaxis: {
				buckets: null
				mode:    "time"
				name:    null
				show:    true
				values: []
			}
			yaxes: [{
				format:  "percentunit"
				label:   null
				logBase: 1
				max:     null
				min:     "0"
				show:    true
			}, {
				format:  "short"
				label:   null
				logBase: 1
				max:     null
				min:     null
				show:    true
			}]
			yaxis: {
				align:      false
				alignLevel: null
			}
		}, {
			aliasColors: {
				max: "dark-yellow"
			}
			bars:        false
			dashLength:  10
			dashes:      false
			datasource:  "$datasource"
			description: "Memory utilisation and limits."
			fieldConfig: {
				defaults: {
					custom: {}
					links: []
				}
				overrides: []
			}
			fill:         0
			fillGradient: 0
			gridPos: {
				h: 8
				w: 6
				x: 12
				y: 24
			}
			hiddenSeries: false
			id:           16
			interval:     "1m"
			legend: {
				avg:     false
				current: false
				max:     false
				min:     false
				show:    true
				total:   false
				values:  false
			}
			lines:     true
			linewidth: 1
			links: []
			maxDataPoints: 250
			nullPointMode: "null"
			options: {
				alertThreshold: true
			}
			percentage:    false
			pluginVersion: "7.4.5"
			pointradius:   2
			points:        false
			renderer:      "flot"
			seriesOverrides: [{
				alias:        "Memory"
				fill:         1
				fillGradient: 5
			}, {
				alias:  "Request"
				color:  "#FF9830"
				dashes: true
			}, {
				alias:  "Limit"
				color:  "#F2495C"
				dashes: true
			}]
			spaceLength: 10
			stack:       false
			steppedLine: false
			targets: [{
				expr:           "avg by (pod) (container_memory_usage_bytes{container=\"cert-manager\"})"
				format:         "time_series"
				hide:           false
				interval:       ""
				intervalFactor: 1
				legendFormat:   "Memory {{pod}}"
				refId:          "A"
			}, {
				expr:           "avg by (pod) (kube_pod_container_resource_limits_memory_bytes{container=\"cert-manager\"})"
				format:         "time_series"
				interval:       ""
				intervalFactor: 1
				legendFormat:   "Limit {{pod}}"
				refId:          "B"
			}, {
				expr:           "avg by (pod) (kube_pod_container_resource_requests_memory_bytes{container=\"cert-manager\"})"
				format:         "time_series"
				interval:       ""
				intervalFactor: 1
				legendFormat:   "Request {{pod}}"
				refId:          "C"
			}]
			thresholds: []
			timeFrom: null
			timeRegions: []
			timeShift: null
			title:     "Memory"
			tooltip: {
				shared:     true
				sort:       0
				value_type: "individual"
			}
			type: "graph"
			xaxis: {
				buckets: null
				mode:    "time"
				name:    null
				show:    true
				values: []
			}
			yaxes: [{
				format:  "bytes"
				label:   null
				logBase: 1
				max:     null
				min:     "0"
				show:    true
			}, {
				format:  "short"
				label:   null
				logBase: 1
				max:     null
				min:     null
				show:    true
			}]
			yaxis: {
				align:      false
				alignLevel: null
			}
		}, {
			aliasColors: {
				max: "dark-yellow"
			}
			bars:        false
			dashLength:  10
			dashes:      false
			datasource:  "$datasource"
			description: "Network ingress/egress."
			fieldConfig: {
				defaults: {
					custom: {}
					links: []
				}
				overrides: []
			}
			fill:         1
			fillGradient: 5
			gridPos: {
				h: 8
				w: 6
				x: 18
				y: 24
			}
			hiddenSeries: false
			id:           18
			interval:     "1m"
			legend: {
				avg:     false
				current: false
				max:     false
				min:     false
				show:    true
				total:   false
				values:  false
			}
			lines:     true
			linewidth: 1
			links: []
			nullPointMode: "null"
			options: {
				alertThreshold: true
			}
			percentage:    false
			pluginVersion: "7.4.5"
			pointradius:   2
			points:        false
			renderer:      "flot"
			seriesOverrides: [{
				alias:     "transmit"
				transform: "negative-Y"
			}]
			spaceLength: 10
			stack:       false
			steppedLine: false
			targets: [{
				expr: """
					avg(
					sum without (interface) (
					    rate(container_network_receive_bytes_total{namespace="cert-manager"}[$__rate_interval])
					)
					)
					"""
				format:         "time_series"
				hide:           false
				interval:       ""
				intervalFactor: 2
				legendFormat:   "receive"
				refId:          "A"
			}, {
				expr: """
					avg(
					sum without (interface) (
					    rate(container_network_transmit_bytes_total{namespace="cert-manager"}[$__rate_interval])
					)
					)
					"""
				format:         "time_series"
				hide:           false
				interval:       ""
				intervalFactor: 2
				legendFormat:   "transmit"
				refId:          "B"
			}]
			thresholds: []
			timeFrom: null
			timeRegions: []
			timeShift: null
			title:     "Network"
			tooltip: {
				shared:     true
				sort:       0
				value_type: "individual"
			}
			type: "graph"
			xaxis: {
				buckets: null
				mode:    "time"
				name:    null
				show:    true
				values: []
			}
			yaxes: [{
				format:  "Bps"
				label:   null
				logBase: 1
				max:     null
				min:     null
				show:    true
			}, {
				format:  "short"
				label:   null
				logBase: 1
				max:     null
				min:     null
				show:    true
			}]
			yaxis: {
				align:      false
				alignLevel: null
			}
		}]
		refresh:       "1m"
		schemaVersion: 27
		style:         "dark"
		tags: ["cert-manager", "infra"]
		templating: {
			list: [{
				current: {
					selected: false
					text:     "prometheus"
					value:    "prometheus"
				}
				description: null
				error:       null
				hide:        0
				includeAll:  false
				label:       "Data Source"
				multi:       false
				name:        "datasource"
				options: []
				query:       "prometheus"
				queryValue:  ""
				refresh:     1
				regex:       ""
				skipUrlSync: false
				type:        "datasource"
			}]
		}
		time: {
			from: "now-24h"
			to:   "now"
		}
		timepicker: {
			refresh_intervals: ["10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"]
		}
		timezone: "browser"
		title:    "Cert Manager"
		uid:      "TvuRo2iMk"
		version:  1
	}
}