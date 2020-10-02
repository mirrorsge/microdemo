package main

import (
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
	s := grpc.NewServer()
	pb.RegisterAlphaServiceServer(s, &handlers.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve")
	}
}
