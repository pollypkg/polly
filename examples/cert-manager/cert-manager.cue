package polly

import "github.com/pollypkg/polly/schema:pollyschema"

// Enforce that the emit value of this file unifies with the Polly schema
pollyschema.PollyPackage

header: {
    name: "cert-manager"
    uri:  "github.com/pollypkg/polly/examples/cert-manager"
    params: {
        certManagerCertExpiryDays: int | *21,
        certManagerJobLabel: string | *"cert-manager",
        // TODO Runbooks are an open question themselves, and not really sure how
        // we'd even think about the interpolation-inside-interpolation here
        // certManagerRunbookURLPattern: "https://gitlab.com/uneeq-oss/cert-manager-mixin/-/blob/master/RUNBOOK.md#%s",
        grafanaExternalUrl: string,
    }
}

prometheusAlerts: v0: {
    CertManagerCertExpirySoon: {
        group: "certificates"
        alert: {
            expr: """
            avg by (exported_namespace, namespace, name) (
            certmanager_certificate_expiration_timestamp_seconds - time()
            ) < (\(header.params.certManagerCertExpiryDays\) * 24 * 3600)
            """
        }
        "for": "1h",
        labels: {
            severity: "warning"
        }
        annotations: {
            summary: "The cert `{{ $labels.name }}` is {{ $value | humanizeDuration }} from expiry, it should have renewed over a week ago.",
            description: "The domain that this cert covers will be unavailable after {{ $value | humanizeDuration }}. Clients using endpoints that this cert protects will start to fail in {{ $value | humanizeDuration }}.",
            // TODO this is totally broken right now because it relies on a
            // hardcoded uid for the particular dashboard. Polly provides the
            // necessary namespacing information such that it should no longer be
            // necessary to sling around uids like this - instead, this should be
            // a reference to the namespaced name of the polly dashboard.
            //
            // That's the ideal, anyway - we'll have to see what we can actually
            // accomplish :)
            dashboard_url: header.params.grafanaExternalUrl + "/d/TvuRo2iMk/cert-manager",
        }
    }
}