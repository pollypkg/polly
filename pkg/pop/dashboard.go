package pop

import (
	"cuelang.org/go/cue"
)

type Dashboard struct {
	v    cue.Value
	name string
}

func (p *Pop) Dashboards() []Dashboard {
	var dbs []Dashboard
	iter, err := p.v.LookupPath(cue.ParsePath("grafanaDashboards.v0")).Fields()
	if err != nil {
		boom(err)
	}
	for iter.Next() {
		db := Dashboard{
			v:    iter.Value(),
			name: iter.Label(),
		}
		dbs = append(dbs, db)
	}

	return dbs
}

func (p *Pop) Dashboard(name string) (*Dashboard, error) {
	v := p.v.LookupPath(cue.ParsePath("grafanaDashboards.v0." + name))
	if v.Err() != nil {
		return nil, v.Err()
	}

	db := Dashboard{
		v: v, name: name,
	}
	return &db, nil
}

// TODO: Implement once we have Go structs
// func (d *Dashboard) Model() (grafana.Dashboard, error) {
// }

// Interface returns the dashboard model as a map[string]interface{} with the
// (weak) guarantee, that it conforms the schema.
// Warning: Will be deprecated once Model() is available
func (d *Dashboard) Interface() (map[string]interface{}, error) {
	var m map[string]interface{}
	if err := d.v.Decode(&m); err != nil {
		return nil, err
	}

	return m, nil
}

func (d *Dashboard) Value() cue.Value {
	return d.v
}

func (d *Dashboard) Name() string {
	return d.name
}
