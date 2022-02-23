package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	*grpc.Server
	dbClient     *mongo.Client
	dbCollection *mongo.Collection
}

func newServer() (*server, error) {
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)

	fmt.Println("Connecting to MongoDB")

	dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	if err := dbClient.Connect(context.Background()); err != nil {
		return nil, err
	}

	dbCollection := dbClient.Database("mydb").Collection("todo")

	reflection.Register(s)

	return &server{
		s,
		dbClient,
		dbCollection,
	}, nil
}

func (s *server) stop() {
	defer s.Stop()

	if err := s.dbClient.Disconnect(context.Background()); err != nil {
		log.Fatalf("faied to disconnect MongoDB %v", err)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Todo service started")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server, err := newServer()
	if err != nil {
		log.Fatalf("failed create server: %v", err)
	}
	defer server.stop()

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
