package main

import (
	"context"
	"errors"
	"log"
	"net"
	"strings"

	pb "github.com/alexuserid/cServ/proto"
	"google.golang.org/grpc"
)

func verify(param string) (string, error) {
	var s string

	if n := len(param); n == 0 || n > 50 {
		return "", errors.New("Incorrect parameter. Wrong size.")
	}

	s = strings.ToLower(param)
	for _, c := range s {
		if c != 'b' && c != 'g' && c != 'r' {
			return "", errors.New("Incorrect parameter. Use only 'r', 'g', 'b'.")
		}
	}
	return s, nil
}

func stones(s string) int32 {
	var c int32
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			c++
		}
	}
	return c
}

type server struct{}

func (s *server) GetNum(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	colors, err := verify(in.Colors)
	if err != nil {
		return nil, err
	}
	actions := stones(colors)
	return &pb.Reply{MovedStones: actions}, nil
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
