package main

import (
	"fmt"
	"log"
	"net"

	pb "sync_tree/api"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedInfoServer
	pb.UnimplementedMarketServer
	pb.UnimplementedUserServer
}

func main() {
	createNewUsers()
	createStartMarket()
	fmt.Sprintln("Server has started! port: 8080 open")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serv := grpc.NewServer()
	pb.RegisterInfoServer(serv, &server{})
	pb.RegisterUserServer(serv, &server{})
	// pb.RegisterMarketServer(serv, &server{})
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
