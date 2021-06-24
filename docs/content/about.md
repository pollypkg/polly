# About

With polly, we're building on the shoulder of giants, or in other words: this
is not the first attempt to make o11y tooling templating work. It all started out in
04/2018 with Frederic Branczyk proposing [Jsonnet Package Management][jpm]
which was further on adopted in several places such as [Prometheus Monitoring
Mixins][prom-mixins] which had the tagline:

> A mixin is a set of Grafana dashboards and Prometheus rules and alerts, 
> packaged together in a reuseable and extensible bundle. Mixins are written 
> in jsonnet, and are typically installed and updated with jsonnet-bundler.

The Mixins, based on jsonnet and maintained alongside the tool turned out to
be a great idea, however, over time certain challenges and limitations became
apparent.

These shortcomings led to a desire to improve on the existing work on Mixins,
resulting in what we now know as polly. Initially simply called Mixins-NG, 
the work around polly kicked off publickly in [April 2021][polly-kickoff], with
Sam Boyer's mail to the newly created Mixins mailing list. 

At that point, Sam already had been working on the concept and core formalization
for several months, broadening the group of folks involved, initially a small 
group of interested folks, a mixture of past and future stakeholders and 
contributors to Mixins.

At this point, polly doesn't have a formal governance body or rule set, but
we do have:

* A guide on how to [contribute][contrib] to polly.
* The commitment that polly and the artefacts in the organization are available
  under Apache License Version 2.0 and will continue to be so.

For any questions or suggestions, we ask you to use the [polly project
discussion][discussion] section in the main repo.

_Sam, Matthias, Michael_

[jpm]: https://docs.google.com/document/d/1czRScSvvOiAJaIjwf3CogOULgQxhY9MkiBKOQI1yR14/
[prom-mixins]: https://monitoring.mixins.dev/
[polly-kickoff]: https://groups.google.com/g/mixins/c/q8B-nWgfO24
[contrib]: https://github.com/pollypkg/polly/blob/main/CONTRIBUTING.md
[discussion]: https://github.com/pollypkg/polly/discussions
