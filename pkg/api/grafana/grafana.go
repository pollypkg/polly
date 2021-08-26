package grafana

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	"github.com/Masterminds/semver/v3"
	grafana "github.com/grafana/grafana-api-golang-client"
)

// Client provides access to the several API layers of Grafana until a proper
// client package is in place.
type Client struct {
	gapi *grafana.Client
	auth Auth
	url  *url.URL
}

// Auth holds authentication information for the Grafana API
type Auth struct {
	Token string
	Basic *url.Userinfo
	OrgID int64
}

// Info holds information about the remote Grafana instance
type Info struct {
	Commit  string
	Version semver.Version
}

// New returns a grafana.Client for the Grafana instance at baseURL
func New(baseURL string, auth Auth) (*Client, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	if auth.OrgID == 0 {
		auth.OrgID = 1
	}

	gapi, err := grafana.New(base.String(), grafana.Config{
		APIKey:    auth.Token,
		BasicAuth: auth.Basic,
		OrgID:     auth.OrgID,
	})
	if err != nil {
		return nil, err
	}

	c := &Client{
		gapi: gapi,
		auth: auth,
		url:  base,
	}
	return c, nil
}

// API returns an actual grafana/grafana-api-golang-client, which implements
// most operations
func (c *Client) API() *grafana.Client {
	return c.gapi
}

// Info returns metadata about the remote Grafana instance
func (c *Client) Info() (*Info, error) {
	u := *c.url
	u.Path = path.Join(u.Path, "api", "health")

	r, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	var info Info
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		return nil, err
	}

	return &info, err
}
