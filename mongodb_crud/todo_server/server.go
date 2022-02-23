package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	*grpc.Server
}

func newServer() *server {
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)

	reflection.Register(s)

	return &server{
		s,
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Todo service started")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := newServer()
	defer server.Stop()

	handler, err := newMongoDBHandler()
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	defer handler.disconnect()

	go func() {
		fmt.Println("starting server...")
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	<-ch

	fmt.Println("stopping server")
}
