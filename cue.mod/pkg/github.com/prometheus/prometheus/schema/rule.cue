package schema

import "github.com/grafana/grafana/cue/scuemata"

Rule: scuemata.#Family & {
    lineages: [
        [
            { // 0.0
                // TODO docs
                expr: string
                // TODO docs
                "for": string
                // TODO docs
                labels: [string]: string
                // TODO docs
                annotations: [string]: string
            }
        ]
    ]
}