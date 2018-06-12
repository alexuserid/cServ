package main

import (
	"context"
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
	r, err := c.GetNum(context.Background(), &pb.Request{Colors: "RRRGG"})
	if err != nil {
		log.Fatalf("c.GetNum: %v", err)
	}
	log.Print("Message sent")
	log.Printf("Answer is: %v", r)
}
