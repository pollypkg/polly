# Usage

Learn about how to use polly.

For now, polly is mostly spec and we're adding support for tooling as we go.

However, because a foundational goal of polly packages is to make observability
universal - that is, avoiding lock-in to any particular deployment/devops
toolchain - the tools provided here may be less featureful or useful 
than tools that fit polly packages into your existing toolchain 
(e.g. Terraform, jsonnet/Tanka, flux/2). We'll link to such tools as they evolve.

We appreciate examples submissions, please follow the following guidelines:

1. Explain what's the system under observation (SUO), for example a Kubernetes
   cluster or a VM or a database.
1. If you show how to migrate from an existing mixin please link to it.

Other questions to consider:

* What kind of parameters do you think it makes sense for your package to take? Why?
* How did you decide whether it was worth making particular queries into reusable signals?
* Which part of the polly spec was most confusing to work with? If you figured it out, what helped you get there?
* Is there anything about observing the SUO that the polly schema did not allow you to express in the pop?

The PR should contain two parts:

* A dedicated directory under the
  [examples](https://github.com/pollypkg/polly/tree/main/examples) directory
  that contains CUE code and any helpers and dependencies.
* A Markdown file under the [docs
  content](https://github.com/pollypkg/polly/tree/main/docs/content), you 
  are encouraged to use [basics.md](https://github.com/pollypkg/polly/blob/main/docs/content/basics.md)
  as a starting point.
