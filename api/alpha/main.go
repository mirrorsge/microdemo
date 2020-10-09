package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/mirrorsge/microdemo/api/alpha/handlers"
	pb "github.com/mirrorsge/microdemo/gateway/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const address = ":9090"

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("failed to listen", address)
	}
	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(),
	)
	pb.RegisterAlphaServiceServer(s, &handlers.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve")
	}
}
