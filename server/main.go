package main

import (
	"context"
	"log"
	"net"

	pb "github.com/solaimanshadin/grpc-go-example/server/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloWorldServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: req.GetName(),
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()

	pb.RegisterHelloWorldServer(s, &server{})
	log.Println("Server running on http://localhost:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error running server %v", err)
	}
}