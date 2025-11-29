package main

import (
	"context"
	"log"
	"net"

	pb "github.com/solaimanshadin/grpc-go-example/server/pb"
	"google.golang.org/grpc"
)

type calcuatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *calcuatorServer) Add(ctx context.Context, req *pb.CalcutorRequest) (*pb.CalcutorResponse, error) {
	sumOfTwo := req.GetNum1() + req.GetNum2()

	return &pb.CalcutorResponse{
		Result: sumOfTwo,
	}, nil
}

func (s *calcuatorServer) Subtract(ctx context.Context, req *pb.CalcutorRequest) (*pb.CalcutorResponse, error) {
	diffrenceOfTwo := req.GetNum1() - req.GetNum2()

	return &pb.CalcutorResponse{
		Result: diffrenceOfTwo,
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()

	pb.RegisterCalculatorServer(s, &calcuatorServer{})
	log.Println("Server running on http://localhost:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error running server %v", err)
	}
}