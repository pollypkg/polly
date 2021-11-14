package polly

import "github.com/pollypkg/polly/schema:pollyschema"

// Enforce that the emit value of this file unifies with the Polly schema
pollyschema.PollyPackage


header: {
	name: "node-exporter"
	uri:  "github.com/pollypkg/polly/examples/node-exporter"
	params: {
		jobval: string | *"node"

		// Select the metrics coming from the node exporter. Note that all
		// the selected metrics are shown stacked on top of each other in
		// the 'USE Method / Cluster' dashboard. Consider disabling that
		// dashboard if mixing up all those metrics in the same dashboard
		// doesn't make sense (e.g. because they are coming from different
		// clusters).
		nodeExporterSelector: ({ job: {
			Value: jobval
			Op: "="
		}} & #SelectorParameterGroup)

		// Select the fstype for filesystem-related queries. If left
		// empty, all filesystems are selected. If you have unusual
		// filesystem you don't want to include in dashboards and
		// alerting, you can exclude them here, e.g. 'fstype!="tmpfs"'.
		fsSelector: #SelectorParameterGroup | *({ fstype: {
			Value: ""
			Op: "!="
		}} & #SelectorParameterGroup)

		// Select the device for disk-related queries. If left empty, all
		// devices are selected. If you have unusual devices you don't
		// want to include in dashboards and alerting, you can exclude
		// them here, e.g. 'device!="tmpfs"'.
		diskDeviceSelector: #SelectorParameterGroup | *({ device: {
			Value: ""
			Op: "!="
		}} & #SelectorParameterGroup)
	}
}

#SelectorParameter: {
	Label: string
	Value: string
	Op: *"=" | "!=" | "=~" | "!=~"
}

#SelectorParameterGroup: [L=string]: #SelectorParameter & {
	Label: L
}

////////////////////////////////
// Imagine that the below are in a CUE file written by the consumer of this pop

header: params: jobval: "whooziwhatsit"
header: params: nodeExporterSelector: {
	job: {
		Value: "node"
		Op: "="
	}
	foo: {
		Value: "bar"
		Op: "="
	}
}