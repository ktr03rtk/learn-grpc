package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ktr03rtk/learn-grpc/mongodb_crud/todopb"
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
	defer lis.Close()

	s := newServer()
	defer s.GracefulStop()

	handler, err := newMongoDBHandler()
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	defer handler.disconnect()

	todopb.RegisterTodoServiceServer(s.Server, handler)

	go func() {
		fmt.Println("starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	<-ch

	fmt.Println("stopping server")
}
