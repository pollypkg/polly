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

type Watcher struct {
	c *fuge.Fuge
	e chan error

	gapi *grafana.Client
	subs map[string]*fuge.Sub
	auth Auth
}

func (c *Client) Watcher() (*Watcher, error) {
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
		e: make(chan error),

		gapi: c.gapi,
		subs: make(map[string]*fuge.Sub),
		auth: c.auth,
	}
	return &w, nil
}

type ChangeHandler func(interface{}) error

func (w *Watcher) Add(uid string, handler ChangeHandler) error {
	channel := fmt.Sprintf("%d/grafana/dashboard/uid/%s", w.auth.OrgID, uid)
	onPub := func(s *centrifuge.Subscription, e centrifuge.PublishEvent) {
		var event struct {
			UID    string
			Action string
			UserID string
		}

		if err := json.Unmarshal(e.Data, &event); err != nil {
			w.e <- err
			return
		}

		d, err := w.gapi.DashboardByUID(uid)
		if err != nil {
			w.e <- err
			return
		}

		if err := handler(d.Model); err != nil {
			w.e <- err
			return
		}
	}

	s, err := w.c.Sub(channel, onPub)
	if err != nil {
		return err
	}

	w.subs[uid] = s
	return nil
}

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
