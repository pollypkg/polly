## Cert Manager

Ported from https://gitlab.com/uneeq-oss/cert-manager-mixin

## Known Issues

* Grafana dashboard doesn't validate, of course, because the schema aren't mature
* Grafana dashboard contains ids that shouldn't be present. See https://github.com/pollypkg/polly/issues/23

## Unconverted things

* There's this [tests.yaml](https://gitlab.com/uneeq-oss/cert-manager-mixin/-/blob/master/tests.yaml) which isn't a standard concept for mixins, and there's no way to pull it over.
* Didn't include runbook URLs, because they're a jsonnet string template within a string template, and because including fully general ones is a bit of an iffy practice, anyway.