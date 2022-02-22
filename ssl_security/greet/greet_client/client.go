package main

import (
	"context"
	"log"

	"greet/greetpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	certFile := "../../ssl/ca.crt"
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		log.Fatal(err)
	}

	opts := grpc.WithTransportCredentials(creds)

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "John",
			LastName:  "Smith",
		},
	}
	res, err := c.UnaryGreet(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
