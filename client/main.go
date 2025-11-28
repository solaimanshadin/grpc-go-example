package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/solaimanshadin/grpc-go-example/client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	name := flag.String("name", "world", "Name to great")
	flag.Parse()
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	defer conn.Close()
	c := pb.NewHelloWorldClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	r, err := c.SayHello(ctx, &pb.HelloRequest{
		Name: *name,
	})

	defer cancel()

	if err != nil {
		log.Fatalf("Could not connect to gRPC server %v", err)
	}

	log.Printf("Hello %s", r.GetMessage())

}
