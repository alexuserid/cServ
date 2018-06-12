package main

import (
	"context"
	"log"
	"net"

	pb "cServ/proto"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GetNum(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{MovedStones: 4}, nil
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("net.Listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStonesServer(s, &server{})
	s.Serve(ln)
}
