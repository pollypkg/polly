# Polly

[Polly](#Naming) is a specification for parameterized packages of observability-related configuration objects - e.g. dashboards, alerts, but also more.

Polly packages are intended to address fundamental challenges involved in keeping observability systems in sync with the systems they're intended to observe:

* Who creates and maintains these objects? (Who knows enough to do it? To do it well? For OSS systems we didn't write ourselves?)
* How can we be sure these objects are the right ones for the versions of our systems?
* How can we define common properties, invariants, or aggregations over these objects across an entire organization? When the organization's telemetry data is heterogenous?

We see the full lifecycle these configuration objects as within scope of polly. More background is in [this doc, defining polly's scope](https://docs.google.com/document/d/1GU0DGy-X6z4FVwbJYPsBKRdqApi2RppW0q2U6YUXOp8), including how and why polly evolved out of lessons learned from [Prometheus monitoring mixins](https://monitoring.mixins.dev/).

This repository contains the polly specification [itself](https://github.com/pollypkg/polly/blob/main/schema/pollypkg.cue), [example](https://github.com/pollypkg/polly/tree/main/examples) polly packages, as well as tools for working with Polly packages (TODO).

## Vision

All systems need observability. That starts with emitting telemetry data. But raw data alone isn’t enough. There’s a set of steps - roughly: collect, transform, interpret, act - through which data passes that are necessary for software to be not merely observable, but _observed_ by real, actual humans. These steps are the collective responsibility of an overall observability platform, driven by many interrelated bits of configuration.

Of course, as software evolves, so too must the configuration that allows it to be observed. By nature, software and its observers are companions. Today’s tooling, however, makes that co-evolution quite friction-ful and error-prone. That’s what polly aims to change - no matter how you ship software, your observability can ride sidecar. We believe that achieving this has the potential to make observability go truly mainstream, much as testing has over the last decade.

## Using Polly

For now, polly is mostly spec! We'll add supporting tooling as we go.

However, because a foundational goal of polly packages is to make observability universal - that is, avoiding lock-in to any particular deployment/devops toolchain - the tools provided here may be less featureful or useful than tools that fit polly packages into your existing toolchain (e.g. [Terraform](https://github.com/hashicorp/terraform), [jsonnet/Tanka](https://github.com/grafana/tanka), [flux](https://github.com/fluxcd/flux)/[2](https://github.com/fluxcd/flux2)). We'll link to such tools as they evolve.

## Naming

The name "Polly" is a pseudo-acronym, loosely derived from (take your pick):

* Parameterized (P) Observability (o11y) configuration packages
* Packages (P) of Observability (o11y) configuration