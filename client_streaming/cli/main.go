package main

import (
	"client/gen/pb"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
)

const (
	fileName   = "cat.jpg"
	bufferSize = 1024
)

func main() {
	fmt.Println("Image upload start")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewImageUploadServiceClient(conn)

	client, err := c.Upload(context.Background())
	if err != nil {
		log.Fatalf("failed to upload: %v", err)
	}

	if err := client.Send(&pb.ImageUploadRequest{
		File: &pb.ImageUploadRequest_FileMeta_{
			FileMeta: &pb.ImageUploadRequest_FileMeta{
				Filename: fileName,
			},
		},
	}); err != nil {
		log.Fatalf("failed to send filename: %v", err)
	}

	fp, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer fp.Close()

	buf := make([]byte, bufferSize)
	i := 0
	for {
		_, err := fp.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("---------------count %+v\n", i)
		i++

		if err := client.Send(&pb.ImageUploadRequest{
			File: &pb.ImageUploadRequest_Data{
				Data: buf,
			},
		}); err != nil {
			log.Fatalf("failed to send data: %v", err)
		}
	}

	res, err := client.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receive data: %v", err)
	}

	fmt.Printf("--------------- %+v\n", res.String())

	fmt.Println("Image upload finished")
}
