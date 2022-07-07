package main

import pb "jimmykiang/grpcCalc/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
