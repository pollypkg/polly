// Package proto defines the gRPC API that the ODE frontend uses to communicate
// with the editing backend (pop edit).
//
// The *.proto files listed below define the actual interfaces. protoc converts
// those into respective definitions and clients:
// - Go: protoc-gen-go and protoc-gen-go-grpc (these must be installed locally)
// - TS: protoc-gen-grpc-web (automatically installed by yarn in node_modules/)
//
// The backend then serves this gRPC API over HTTP (using grpc-web), which the
// React (TypeScript) application connects to.
package proto

//go:generate ./gen.sh dashboard.proto
