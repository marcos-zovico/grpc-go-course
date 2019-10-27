package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/msouza/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {

	fmt.Println("CAlculator Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Faliled to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (*server) Sun(ctx context.Context, req *calculatorpb.CalcRequest) (*calculatorpb.CalcResponse, error) {
	fmt.Printf("Sun function was invoked with %v\n", req)
	result := req.GetOperation().GetNumberOne() + req.GetOperation().GetNumberTwo()

	response := &calculatorpb.CalcResponse{
		Result: result,
	}
	return response, nil
}

func (*server) Multiply(ctx context.Context, req *calculatorpb.CalcRequest) (*calculatorpb.CalcResponse, error) {
	fmt.Printf("Multiply function was invoked with %v\n", req)

	result := req.GetOperation().GetNumberOne() * req.GetOperation().GetNumberTwo()

	response := &calculatorpb.CalcResponse{
		Result: result,
	}
	return response, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Received PrimeNumberDecomposition RPC: %v\n", req)

	number := req.GetNumber()
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v\n", divisor)
		}
	}
	return nil
}