package main

import (
	"context"
	"log"
	"time"

	pb "github.com/solaimanshadin/grpc-go-example/client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	var  num1 float64 = 1
	var  num2 float64 = 1
	r, err := c.Add(ctx, &pb.CalcutorRequest{
		Num1: num1,
		Num2: num2,
	})

	defer cancel()

	if err != nil {
		log.Fatalf("Could not connect to gRPC server %v", err)
	}

	log.Printf("Sum of %f and %f is: %f", num1, num2, r.GetResult())
}
