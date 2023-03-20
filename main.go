package main

import (
	"context"
	"fmt"
	"log"
	"net"

	calculatorPB "gogrpc/calculator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	calculatorPB.UnimplementedCalculatorServiceServer
}

func (*Server) Sum(ctx context.Context, req *calculatorPB.CalculatorRequest) (*calculatorPB.CalculatorResponse, error) {
	fmt.Printf("Sum function is invoked with %v \n", req)

	a := req.GetA()
	b := req.GetB()

	res := &calculatorPB.CalculatorResponse{
		Result: a + b,
	}

	return res, nil
}

func main() {
	fmt.Println("starting gRPC server...")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	calculatorPB.RegisterCalculatorServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
