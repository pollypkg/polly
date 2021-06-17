# Basic usage

We have a very basic example in [examples/basic/](https://github.com/pollypkg/polly/blob/main/examples/basic) that you can use to explore polly packages.
This example contains multiple signals (in this case PromQL queries), a dataface, and in the end generates a Grafana dashboard with those lower level constructs.

### Evaluating

Because the example comes with parameterized parameters (params) we cannot straight export JSON or  YAML, but instead can evaluate the example. This means that CUE will evaluate as much as possible of the end result, still leaving a few places without the final parameters/strings.

From the root of the repository you can evaluate by running:

```bash
cue eval ./examples/basic
```

*Note: If any of the constraints aren't fulfilled this step will fail and let you know that something is wrong.*

### Exporting

To export you need to insert actual parameters to the polly package for it to be able to actually export real YAML. Let's give the example some concrete params:

_Note: directly modifying a polly package is NOT how we expect polly to be used in practice. Rather, you'll inject these parameter values via the tool that consumes the polly package._

```diff
 		{
 			name: "NumCpu"
 			lang: "promql"
-			params: {job: string, instance: string}
+			params: {job: "node", instance: "localhost:9100"}
 			query: "count without (cpu) (count without (mode) (node_cpu_seconds_total{job=\"\(params.job)\", instance=\"\(params.instance)\"}))"
 		},
 		// Amount of memory currently in use
 		{
 			name: "MemoryUtilization"
 			lang: "promql"
-			params: {job: string, instance: string}
+			params: {job: "node", instance: "localhost:9100"}
 			query: "1 - (node_memory_MemAvailable_bytes{job=\"\(params.job)\", instance=\"\(params.instance)\"} / node_memory_MemTotal_bytes{job=\"\(params.job)\", instance=\"\(params.instance)\"})"
 		},
 		// One minute rate of major page faults
 		{
 			name: "VmstatPGMajFault"
 			lang: "promql"
-			params: {job: string, instance: string}
+			params: {job: "node", instance: "localhost:9100"}
 			query: "rate(node_vmstat_pgmajfault{job=\"\(params.job)\", instance=\"\(params.instance)\"}[1m])"
 		},
 	]

```

With that change, we can now run `cue export ./examples/basic ` and will get a JSON output.
Run `cue export --out yaml ./examples/basic` if you prefer YAML instead.

