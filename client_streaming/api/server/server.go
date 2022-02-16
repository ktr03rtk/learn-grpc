package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"image.upload/gen/pb"
	"image.upload/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterImageUploadServiceServer(server, handler.NewImageUploadHandler())
	reflection.Register(server)

	go func() {
		log.Fatalf("start gRPC server port: %v", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server..")
	server.GracefulStop()
}
