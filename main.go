package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/devholic77/grpc-echo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.EchoServer
}

func (e *server) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Message: fmt.Sprintf("echo:%s, from localhost:%s", req.Message, port),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
