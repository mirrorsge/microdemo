package handlers

import (
	"context"
	pb "github.com/mirrorsge/microdemo/gateway/proto"
)

type Server struct {
	pb.UnimplementedAlphaServiceServer
}

func (s *Server) Hello(context.Context, *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Name: "lijianjun",
	}, nil
}
