package edit

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	gapi "github.com/grafana/grafana-api-golang-client"

	"cuelang.org/go/cue/format"
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

		inEdit: make(map[string]*pop.Dashboard),
		api:    opts.Client,
		watch:  w,
	}

	return &g, nil
}

type Grafana struct {
	p pop.Pop

	inEdit map[string]*pop.Dashboard

	api   *grafana.Client
	watch *grafana.Watcher
}

func (c Grafana) Add(name string) error {
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

	originalUID := i["uid"]
	editUID := dashboardID(name)
	i["uid"] = editUID
	i["id"] = nil

	_ = c.api.API().DeleteDashboardByUID(editUID)
	_, err = c.api.API().NewDashboard(gapi.Dashboard{
		Model: i,
	})
	if err != nil {
		return err
	}

	c.inEdit[editUID] = d

	err = c.watch.Add(editUID, func(i map[string]interface{}) error {
		i["uid"] = originalUID
		delete(i, "id")

		model := map[string]interface{}{
			"grafanaDashboards": map[string]interface{}{
				"v0": map[string]interface{}{
					name: i,
				},
			},
		}

		data, err := json.MarshalIndent(model, "", "  ")
		if err != nil {
			panic(err)
		}

		fmted, err := format.Source(data, format.Simplify())
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

func (c Grafana) Close() error {
	var outErr error
	for uid := range c.inEdit {
		err := c.api.API().DeleteDashboardByUID(uid)
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
