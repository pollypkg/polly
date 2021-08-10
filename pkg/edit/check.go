package edit

import (
	"fmt"
	"path/filepath"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"cuelang.org/go/cue/token"
	"github.com/pollypkg/polly/pkg/pop"
)

// Check reports whether this pop can be edited
func Check(p pop.Pop) error {
	dashboards := p.Value().LookupPath(cue.ParsePath("grafanaDashboards.v0"))
	i, err := dashboards.Fields()
	if err != nil {
		return err
	}
	for i.Next() {
		f, err := ValueFile(i.Value())
		if err != nil {
			return err
		}

		if err := containsOnly(f, i.Value().Path()); err != nil {
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

// ValueFile reports the file a value originates from.
// It expects values to originate from _exactly one_ file.
// The polly schema is ignored
func ValueFile(v cue.Value) (string, error) {
	split := v.Split()

	pos := make([]token.Pos, len(split))
	for i, s := range split {
		pos[i] = s.Pos()
	}

	if len(split) == 0 {
		panic("shouldn't happen. figure out why")
	}

	if len(split) == 1 {
		return pos[0].Filename(), nil
	}

	if len(split) > 2 {
		return "", ErrMultipleFiles{pos: pos, name: v.Path().String()}
	}

	var file string
	for _, p := range pos {
		// don't count our schema
		if isSchema(p.Filename()) {
			continue
		}

		if file != "" {
			return "", ErrMultipleFiles{pos: pos, name: v.Path().String()}
		}

		file = p.Filename()
	}

	if file == "" {
		panic("shouldn't happen. figure out why")
	}

	return file, nil
}

type ErrMultipleFiles struct {
	pos  []token.Pos
	name string
}

func (e ErrMultipleFiles) Error() string {
	s := fmt.Sprintf("The dashboard '%s' originates from more than one non-schema file:\n", e.name)

	for _, p := range e.pos {
		s += "  - " + p.String()
		if isSchema(p.Filename()) {
			s += " (schema)"
		}
		s += "\n"
	}

	s += "Editing requires dashboards to be in their separate files"
	return s
}

func isSchema(f string) bool {
	// TODO: this is a very naive and weak assumption. find a better one
	return strings.HasSuffix(filepath.ToSlash(f), "polly/schema/pollypkg.cue")
}
