//go:build !test
// +build !test

package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "jimmykiang/grpcCalc/calculator/proto"
	"log"
	"net"
)

//var addr string = "0.0.0.0:50051"

// for docker
var addr string = ":50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
