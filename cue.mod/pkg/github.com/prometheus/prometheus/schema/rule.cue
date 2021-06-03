package schema

import "github.com/grafana/grafana/cue/scuemata"

Rule: scuemata.#Family & {
    lineages: [
        [
            {
                placeholder: string
            }
        ]
    ]
}