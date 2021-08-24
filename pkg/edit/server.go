package edit

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/pollypkg/polly/pkg/coord"
	"github.com/pollypkg/polly/pkg/pop"
)

// Server provides edit operations over a HTTP interface
type Server struct {
	p pop.Pop

	g *Grafana
}

// HTTPServer returns a http.Handler serving the primary api
func HTTPServer(ctx context.Context, p pop.Pop, opts Opts) (http.Handler, error) {
	g, err := Edit(p, opts)
	if err != nil {
		return nil, err
	}

	s := Server{
		p: p,
		g: g,
	}

	// cleanup once context canceled
	coord.Finally(ctx, func() {
		if err := s.Close(); err != nil {
			log.Println(err)
		}
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/dashboards", s.Dashboards)
	mux.HandleFunc("/edit", s.Edit)

	// TODO: properly handle CORS
	corsDisabled := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		mux.ServeHTTP(w, r)
	})

	return corsDisabled, nil
}

// DashboardMeta holds basic information about a Dashboard tailored towards
// displaying in the UI
type DashboardMeta struct {
	File string `json:"file"`
	Name string `json:"name"`

	Title string `json:"title"`
	UID   string `json:"uid"`
	Desc  string `json:"desc"`
}

// List Dashbaords
// GET /dashboards
func (s Server) Dashboards(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "/dashboards only supports GET", http.StatusBadRequest)
		return
	}

	var meta []DashboardMeta
	for _, d := range s.p.Dashboards() {
		var m DashboardMeta
		if err := d.Value().Decode(&m); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, err := File(d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		m.Name = d.Name()
		m.File = file
		meta = append(meta, m)
	}

	data, err := json.Marshal(meta)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

// Edit a dashboard
// POST /edit?dashboard=<name>
func (s Server) Edit(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("dashboard")
	if name == "" {
		http.Error(w, "dashboard url parameter required", http.StatusBadRequest)
		return
	}

	if err := s.g.Add(name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "http://localhost:3000/d/"+s.g.EditUID(name), http.StatusTemporaryRedirect)
}

// Close the server and cleanup any leftovers
func (s Server) Close() error {
	return s.g.Close()
}
