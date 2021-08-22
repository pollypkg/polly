package edit

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"github.com/pollypkg/polly/pkg/pop"
)

// Check reports whether this pop can be edited
func Check(p pop.Pop) error {
	for _, d := range p.Dashboards() {
		f, err := File(d)
		if err != nil {
			return err
		}

		if err := containsOnly(f, d.Value().Path()); err != nil {
			return err
		}
	}

	return nil
}

func containsOnly(file string, path cue.Path) error {
	inst := load.Instances([]string{file}, nil)[0]

	ctx := cuecontext.New()
	top := ctx.BuildInstance(inst)
	if top.Err() != nil {
		return top.Err()
	}

	db := top.LookupPath(cue.ParsePath("grafanaDashboards.v0"))
	if db.Err() != nil {
		return db.Err()
	}

	for _, v := range []cue.Value{top, db} {
		it, err := v.Fields()
		if err != nil {
			return err
		}

		var fields []cue.Value
		for it.Next() {
			fields = append(fields, it.Value())
		}

		if len(fields) > 1 {
			return ErrorNotSole{
				name:   path.String(),
				file:   file,
				fields: fields,
			}
		}
	}

	return nil
}

type ErrorNotSole struct {
	name   string
	file   string
	fields []cue.Value
}

func (e ErrorNotSole) Error() string {
	s := fmt.Sprintf("Dashboard '%s' is not sole resource in file:\n", e.name)
	for _, f := range e.fields {
		s += fmt.Sprintf("  - %s", f.Path())
		for _, v := range f.Split() {
			if v.Pos().Filename() == e.file {
				s += fmt.Sprintf(" (%s:%d:%d)", v.Pos().Filename(), v.Pos().Line(), v.Pos().Column())
			}
		}
		s += "\n"
	}
	s += "Editing requires dashboards to be in their separate files"
	return s
}
