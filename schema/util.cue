package pollyschema

import (
	"github.com/grafana/grafana/cue/scuemata"
)

_latest: {
	arg: scuemata.#Lineage
	out: arg[len(arg)-1]
}
