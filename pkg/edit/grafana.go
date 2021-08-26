package edit

import (
	"context"
	"fmt"

	"github.com/pollypkg/polly/pkg/edit/proto"
)

// compile time check that *DashboardService implements the proto server
var _ proto.DashboardServiceServer = &DashboardService{}

// DashboardService provides edit capabilities for Grafana Dashboards
type DashboardService struct {
	g *Grafana

	proto.UnimplementedDashboardServiceServer
}

// List of dashboard metadata
func (s *DashboardService) List(ctx context.Context, r *proto.ListDashboardsRequest) (*proto.ListDashboardResponse, error) {
	meta := []*proto.DashboardMeta{}

	for _, d := range s.g.p.Dashboards() {
		var m proto.DashboardMeta
		if err := d.Value().Decode(&m); err != nil {
			return nil, err
		}

		file, err := File(d)
		if err != nil {
			return nil, err
		}

		m.Name = d.Name()
		m.File = file
		meta = append(meta, &m)
	}

	res := proto.ListDashboardResponse{
		Dashboards: meta,
	}
	return &res, nil
}

// Edit the dashboard specified by it's CUE name
func (s *DashboardService) Edit(ctx context.Context, r *proto.EditDashboardRequest) (*proto.EditDashboardResponse, error) {
	if r.Name == "" {
		return nil, fmt.Errorf("dashboard name must not be empty")
	}

	if err := s.g.Add(r.Name); err != nil {
		return nil, err
	}

	res := proto.EditDashboardResponse{
		EditURL: "http://localhost:3000/d/" + s.g.EditUID(r.Name),
	}
	return &res, nil
}

func (s *DashboardService) Close() error {
	return s.g.Close()
}
