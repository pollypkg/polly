package edit

import (
	"fmt"
	"path/filepath"
	"strings"

	"cuelang.org/go/cue/token"
	"github.com/pollypkg/polly/pkg/pop"
)

// File reports the file a dashboard originates from.
// It expects values to originate from _exactly one_ file.
// The polly schema is ignored in that assumption.
func File(d pop.Dashboard) (string, error) {
	v := d.Value()
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

// ErrMultipleFiles happens when the underlying cue.Value of the dashboard was
// constructed from more than one non-schema file.
// Schema files are determined using isSchema().
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
