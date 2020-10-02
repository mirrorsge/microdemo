package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "github.com/mirrorsge/microdemo/gateway/proto"
	"google.golang.org/grpc"
	"net/http"
)

var grpcServerEndpoint = flag.String("grpc-server-endpoint", ":9090", "gRPC server endpoint")

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := gw.RegisterAlphaServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
