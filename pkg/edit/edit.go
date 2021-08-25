package edit

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	gapi "github.com/grafana/grafana-api-golang-client"

	"cuelang.org/go/cue/format"
	"cuelang.org/go/cue/load"
	"github.com/pollypkg/polly/pkg/api/grafana"
	"github.com/pollypkg/polly/pkg/pop"
)

type Opts struct {
	Client *grafana.Client
}

func Edit(p pop.Pop, opts Opts) (*Grafana, error) {
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

	return &g, nil
}

type Grafana struct {
	p pop.Pop

	inEdit map[string]string

	api   *gapi.Client
	watch *grafana.Watcher
}

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

	originalUID := i["uid"]
	editUID := dashboardID(name)
	i["uid"] = editUID
	i["id"] = nil

	_ = c.api.DeleteDashboardByUID(editUID)
	_, err = c.api.NewDashboard(gapi.Dashboard{
		Model: i,
	})
	if err != nil {
		return err
	}

	c.inEdit[name] = editUID

	err = c.watch.Add(editUID, func(upd map[string]interface{}) error {
		upd["uid"] = originalUID
		delete(upd, "id")
		delete(upd, "version")
		trim(upd)

		model := map[string]interface{}{
			"grafanaDashboards": map[string]interface{}{
				"v0": map[string]interface{}{
					name: upd,
				},
			},
		}

		data, err := json.MarshalIndent(model, "", "  ")
		if err != nil {
			panic(err)
		}

		trimmed := strings.Trim(string(data), "{}")

		pkged := fmt.Sprintf("package %s\n%s", cuePkg, trimmed)

		fmted, err := format.Source([]byte(pkged), format.Simplify())
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile(file, fmted, 0744); err != nil {
			panic(err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (c Grafana) EditUID(name string) string {
	return c.inEdit[name]
}

func (c Grafana) Close() error {
	var outErr error
	for _, uid := range c.inEdit {
		err := c.api.DeleteDashboardByUID(uid)
		if outErr == nil {
			outErr = err
		}
	}

	return outErr
}

func dashboardID(name string) string {
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

func cuePackage(file string) (string, error) {
	inst := load.Instances([]string{file}, nil)
	i := inst[0]
	if i.Err != nil {
		return "", i.Err
	}

	return i.PkgName, nil
}

func trim(i interface{}) {
	switch i := i.(type) {
	case map[string]interface{}:
		for k, v := range i {
			if v == nil {
				delete(i, k)
				continue
			}

			trim(v)
		}
	case []interface{}:
		for _, v := range i {
			trim(v)
		}
	}
}
