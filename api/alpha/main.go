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
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("failed to listen", address)
	}

	s := grpc.NewServer(
		//流式中间件，暂不使用
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer()),
		//一元中间件
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer()),
	)
	pb.RegisterAlphaServiceServer(s, &handlers.Server{})
	log.Println("rpc服务已经开启")
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve")
	}
}
