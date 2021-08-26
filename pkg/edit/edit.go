package edit

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	gapi "github.com/grafana/grafana-api-golang-client"

	"cuelang.org/go/cue/format"
	"cuelang.org/go/cue/load"
	"github.com/pollypkg/polly/pkg/api/grafana"
	"github.com/pollypkg/polly/pkg/pop"
)

// Opts holds optional properties for Edit that either have sensible defaults or
// are not generally required.
type Opts struct {
	Client *grafana.Client
}

// Edit opens a new editing session on the given Polly package
// Make sure to always close the returned Editor once done editing.
func Edit(p pop.Pop, opts Opts) (*Editor, error) {
	w, err := opts.Client.NewWatcher()
	if err != nil {
		return nil, err
	}

	g := Grafana{
		p: p,

		inEdit: make(map[string]string),
		api:    opts.Client.API(),
		watch:  w,
	}

	return &Editor{Grafana: g}, nil
}

// Editor bundles editing capabilites
type Editor struct {
	Grafana Grafana
}

// Grafana edits Grafana dashboards.
// This may be expanded in the future to other Grafana types
type Grafana struct {
	p pop.Pop

	inEdit map[string]string

	api   *gapi.Client
	watch *grafana.Watcher
}

// Add starts interactive editing for the dashboard pointed to by name:
// - Create the dashboard using a temporary UID at the Grafana instance
// - Subscribe to change events, write back to disk when those occur
func (c Grafana) Add(name string) error {
	if _, ok := c.inEdit[name]; ok {
		return nil
	}

	d, err := c.p.Dashboard(name)
	if err != nil {
		return err
	}

	i, err := d.Interface()
	if err != nil {
		return err
	}

	file, err := File(*d)
	if err != nil {
		return err
	}

	cuePkg, err := cuePackage(file)
	if err != nil {
		return err
	}

	// TODO(sh0rez): this is hacky, better formalize (also the undoing below)
	originalUID := i["uid"]
	editUID := genEditUID(name)
	i["uid"] = editUID
	i["id"] = nil

	_ = c.api.DeleteDashboardByUID(editUID)
	if _, err = c.api.NewDashboard(gapi.Dashboard{Model: i}); err != nil {
		return fmt.Errorf("Failed to create temporary dashboard '%s' in Grafana: %w", editUID, err)
	}

	c.inEdit[name] = editUID

	err = c.watch.Add(editUID, func(upd map[string]interface{}, err error) {
		if err != nil {
			log.Printf("Error: Failed receiving update event for '%s': %s", editUID, err)
			return
		}

		upd["uid"] = originalUID
		delete(upd, "id")
		delete(upd, "version")
		Trim(upd)

		model := map[string]interface{}{
			"grafanaDashboards": map[string]interface{}{
				"v0": map[string]interface{}{
					name: upd,
				},
			},
		}

		data, err := json.MarshalIndent(model, "", "  ")
		if err != nil {
			// failure here suggests programming mistakes, as model must be
			// serializable
			panic(err)
		}

		trimmed := strings.Trim(string(data), "{}")

		pkged := fmt.Sprintf("package %s\n%s", cuePkg, trimmed)

		fmted, err := format.Source([]byte(pkged), format.Simplify())
		if err != nil {
			// must not fail, as JSON is valid CUE by definition
			panic(err)
		}

		if err := ioutil.WriteFile(file, fmted, 0744); err != nil {
			log.Printf("Error: Failed to write updated '%s' to disk: %s", file, err)
		}
	})

	return err
}

// EditUID returns the uid of the current editing session of the dashboard
// pointed to by name.
// Returns the empty string if no edit is progress / the dashboard does not exist.
func (c Grafana) EditUID(name string) string {
	return c.inEdit[name]
}

// Close removes all temporary editing dashboards from the Grafana instance and
// stops all update subscriptions.
func (c Grafana) Close() error {
	var lastErr error
	for _, uid := range c.inEdit {
		if err := c.api.DeleteDashboardByUID(uid); err != nil {
			lastErr = err
		}
	}

	if lastErr != nil {
		return lastErr
	}

	return c.watch.Close()
}

// genEditUID generates a editing UID for the dashboard of given name
func genEditUID(name string) string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	if len(hostname) > 11 {
		hostname = hostname[:11]
	}

	hash := base64.RawStdEncoding.EncodeToString(sha256.New().Sum([]byte(name)))[:28]
	id := fmt.Sprintf("pop-%s-%s", hostname, hash)
	if len(id) != 40 {
		panic(len(id))
	}

	return id
}

// cuePackage returns the package used by file.
// TODO: there must be a more efficient way for this (lexical analysis?)
func cuePackage(file string) (string, error) {
	inst := load.Instances([]string{file}, nil)
	i := inst[0]
	if i.Err != nil {
		return "", i.Err
	}

	return i.PkgName, nil
}

// Trim removes unwanted fields from the dashboard model in-place.
// TODO(sh0rez): merge with removing version, id, etc
// TODO: use schema to trim defaults
func Trim(i interface{}) {
	switch i := i.(type) {
	case map[string]interface{}:
		for k, v := range i {
			if v == nil {
				delete(i, k)
				continue
			}

			Trim(v)
		}
	case []interface{}:
		for _, v := range i {
			Trim(v)
		}
	}
}
