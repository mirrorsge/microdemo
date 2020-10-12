package handlers

import (
	"context"
	pb "github.com/mirrorsge/microdemo/gateway/proto"
)

type Server struct {
	pb.UnimplementedAlphaServiceServer
}

func (s *Server) Hello(ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &pb.Response{
		Name: req.Name,
	}, nil
}

func (s *Server) Goodbye(ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &pb.Response{
		Name: req.Name,
	}, nil
}
