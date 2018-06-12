package main

import (
	"context"
	"fmt"
	"log"

	pb "cServ/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial: %v", err)
	}

	c := pb.NewStonesClient(conn)

	var stones string
	fmt.Scan(&stones)
	r, err := c.GetNum(context.Background(), &pb.Request{Colors: stones})
	if err != nil {
		log.Printf("c.GetNum: %v", err)
		return
	}
	log.Printf("Stones to move: %v", r.MovedStones)
}
