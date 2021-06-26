package main

import (
	"fmt"
	"log"
	"net"

	pb "sync_tree/api"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSyncTreeServer
}

func main() {
	fmt.Println("the game goes on")
	setStartBalance()
	createStartMarket()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSyncTreeServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
