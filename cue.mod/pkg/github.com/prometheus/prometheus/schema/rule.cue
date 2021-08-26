package schema

import "github.com/grafana/grafana/cue/scuemata"

Rule: scuemata.#Family & {
	lineages: [
		[
			{// 0.0
				// The name of the time series to output to. Must be a valid metric name.
				record: string

				// The PromQL expression to evaluate. Every evaluation cycle this is
				// evaluated at the current time, and the result recorded as a new set of
				// time series with the metric name as given by 'record'.
				expr: string

				// Labels to add or overwrite before storing the result.
				labels: [string]: string
			},
		],
	]
}
