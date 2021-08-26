package grafana

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/Masterminds/semver/v3"
	"github.com/centrifugal/centrifuge-go"
	grafana "github.com/grafana/grafana-api-golang-client"
	"github.com/pollypkg/polly/pkg/api/fuge"
)

// Watcher listens for dashboard change events and dispatches registered actions
// when they happen
type Watcher struct {
	c *fuge.Fuge

	gapi *grafana.Client
	subs map[string]*fuge.Sub
	auth Auth
}

// NewWatcher returns a new watcher already connected to Grafana's centrifuge
// API. Use Add() to subscribe to the actual dashboard change events.
func (c *Client) NewWatcher() (*Watcher, error) {
	i, err := c.Info()
	if err != nil {
		return nil, err
	}

	if i.Version.LessThan(semver.MustParse("8.0.0")) {
		return nil, fmt.Errorf("live changes requires at least Grafana 8, yours is on %s however", i.Version)
	}

	wsURL := *c.url
	wsURL.Scheme = "ws"
	wsURL.Path = path.Join(wsURL.Path, "api", "live", "ws")
	q := wsURL.Query()
	q.Set("format", "json")
	wsURL.RawQuery = q.Encode()

	f := fuge.New(wsURL.String(), centrifuge.DefaultConfig())
	f.Token(c.auth.Token)
	if err := f.Connect(); err != nil {
		return nil, err
	}

	w := Watcher{
		c: f,

		gapi: c.gapi,
		subs: make(map[string]*fuge.Sub),
		auth: c.auth,
	}
	return &w, nil
}

// ChangeHandler is invoked when a dashboard change event occurs. It is passed
// either the dashboard model, or an error that occured.
type ChangeHandler func(map[string]interface{}, error)

// Add susbcribes to change events for the dashboard of given UID and invokes
// handler when those occur.
func (w *Watcher) Add(uid string, handler ChangeHandler) error {
	channel := fmt.Sprintf("%d/grafana/dashboard/uid/%s", w.auth.OrgID, uid)
	onPub := func(s *centrifuge.Subscription, e centrifuge.PublishEvent) {
		var event struct {
			UID    string
			Action string
			UserID string
		}

		if err := json.Unmarshal(e.Data, &event); err != nil {
			handler(nil, err)
			return
		}

		d, err := w.gapi.DashboardByUID(uid)
		if err != nil {
			handler(nil, err)
			return
		}

		handler(d.Model, nil)
	}

	s, err := w.c.Sub(channel, onPub)
	if err != nil {
		return err
	}

	w.subs[uid] = s
	return nil
}

// Del removes the subscription for the dashboard of given UID if there is one.
func (w *Watcher) Del(uid string) {
	s, ok := w.subs[uid]
	if !ok {
		return
	}

	s.Close()
	delete(w.subs, uid)
}

func (w *Watcher) Close() error {
	var lastErr error
	for _, s := range w.subs {
		if err := s.Close(); err != nil {
			lastErr = err
		}
	}

	return lastErr
}
