package edit

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/pollypkg/polly/pkg/coord"
	"github.com/pollypkg/polly/pkg/edit/proto"
	"github.com/pollypkg/polly/pkg/pop"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// HTTPHandler returns a http.Handler serving the primary gRPC-web API and
// user-interface.
func HTTPHandler(ctx context.Context, p pop.Pop, opts Opts) (http.Handler, error) {
	e, err := Edit(p, opts)
	if err != nil {
		return nil, err
	}

	dashSrv := &DashboardService{g: &e.Grafana}
	coord.Finally(ctx, func() {
		if err := dashSrv.Close(); err != nil {
			log.Println(err)
		}
	})

	// real grpc
	grpcServer := grpc.NewServer()
	proto.RegisterDashboardServiceServer(grpcServer, dashSrv)
	reflection.Register(grpcServer)

	// grpc-web
	webServer := http.StripPrefix("/rpc/v1", grpcweb.WrapServer(grpcServer))

	mux := http.NewServeMux()
	mux.HandleFunc("/rpc/v1", func(w http.ResponseWriter, r *http.Request) {
		// TODO: properly handle CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		webServer.ServeHTTP(w, r)
	})

	srv := Server{
		grpc: grpcServer,
		http: mux,
	}

	return &srv, nil
}

// Server holds an implementation of the edit gRPC API.
// It's ServeHTTP capabilities serve a grpc-web endpoint at /rcp/v1,
// while the ListenGRPC method starts a regular gRPC Server.
type Server struct {
	grpc *grpc.Server
	http *http.ServeMux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.http.ServeHTTP(w, r)
}

func (s *Server) ListenGRPC(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return s.grpc.Serve(lis)
}
