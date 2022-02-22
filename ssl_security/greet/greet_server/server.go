package main

import (
	"context"
	"fmt"
	"greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

func (*server) UnaryGreet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("UnaryGreet %v\n", req.String())
	firstName := req.GetGreeting().GetFirstName()
	result := "hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	certFile := "../../ssl/server.crt"
	keyFile := "../../ssl/server.pem"
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	opts := grpc.Creds(creds)
	s := grpc.NewServer(opts)
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
