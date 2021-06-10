# Welcome!

Polly is a specification for parameterized packages of observability-related
configuration objects such as dashboards, alerts and [more][targets].

Wit polly packages we intend to address the fundamental challenges involved
in keeping observability systems in sync with the systems they're intended to
observe, addressing question including but not limited to:

* Who creates and maintains these objects? Who knows enough to do it? 
  To do it well? For OSS systems we didn't write ourselves?
* How can we be sure these objects are the right ones for the versions of our
  systems?
* How can we define common properties, invariants, or aggregations over these 
  objects across an entire organization? When the organization's telemetry data
  is heterogenous?
  
We see the full [lifecycle](lifecycle.md) these configuration objects as within
scope of polly.

In our repo you will find the [polly specification][spec] and the rest of the
documentation is on this site.

And now: go and give it a try by visiting the [usage](usage.md) section or
check out the [community](community.md) section to learn more about he vision
and how to start contributing.


[spec]: https://github.com/pollypkg/polly/blob/main/schema/pollypkg.cue
[targets]: https://docs.google.com/document/d/1naFdxoO_vAo3GbCM4CpWYwaONxB3c0ar-RMWzqszAsw/
