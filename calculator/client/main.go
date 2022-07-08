package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "jimmykiang/grpcCalc/calculator/proto"
)

// for docker
var addr string = "grpcserver:50051"

//var addr string = "127.0.0.1:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)

	doSum(c)
	// doPrimes(c)
	// doAvg(c)
	// doMax(c)
	// doSqrt(c, 10)
	// doSqrt(c, -2)
}
