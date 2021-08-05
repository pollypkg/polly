package pop

import (
	_ "embed"
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/errors"
	"cuelang.org/go/cue/load"
)

// Pop represents a polly package
type Pop struct {
	v   cue.Value
	ctx *cue.Context
}

// New constructs a pop from a cue.Value
func New(v cue.Value) *Pop {
	return &Pop{v: v, ctx: v.Context()}
}

// Load a polly package from disk
func Load(path string) (*Pop, error) {
	inst := load.Instances([]string{path}, nil)
	if len(inst) != 1 {
		return nil, fmt.Errorf("polly requires exactly one instance. Found %d at '%s'", len(inst), path)
	}

	ctx := cuecontext.New()
	v := ctx.BuildInstance(inst[0])
	if err := v.Err(); err != nil {
		return nil, err
	}

	if err := v.Validate(); err != nil {
		return nil, err
	}

	return New(v), nil
}

// Value returns the underlying cue.Value
func (p Pop) Value() cue.Value {
	return p.v
}

// special CUE file to convert polly packages to mixin format
//go:embed mix.cue
var mixCue string

// Mixin converts the polly package to mixin compatible format
func (p Pop) Mixin() (*Mixin, error) {
	mixer := p.v.Context().CompileString(mixCue,
		cue.Filename("<polly/mix.cue>"),
		cue.Scope(p.v),
	)
	if mixer.Err() != nil {
		panic(fmt.Errorf("failed loading internal mix.cue! Please raise an issue. Error:\n%s", errors.Details(mixer.Err(), nil)))
	}

	mixed := mixer.Unify(p.v)
	if mixed.Err() != nil {
		return nil, mixed.Err()
	}

	mixin := mixer.LookupPath(cue.ParsePath("mixin"))

	m := Mixin{
		GrafanaDashboards: make(map[string]string),
		PrometheusRules:   make([]interface{}, 0),
		PrometheusAlerts:  make([]interface{}, 0),
	}

	if err := mixin.Decode(&m); err != nil {
		return nil, fmt.Errorf("decoding CUE into struct: %w", err)
	}

	return &m, nil
}

// Mixin compatible format representation
type Mixin struct {
	GrafanaDashboards map[string]string `json:"grafanaDashboards"`
	PrometheusRules   []interface{}     `json:"prometheusRules"`
	PrometheusAlerts  []interface{}     `json:"prometheusAlerts"`
}
