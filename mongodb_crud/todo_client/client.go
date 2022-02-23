package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/ktr03rtk/learn-grpc/mongodb_crud/todopb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("TODO client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("failed to connect %v", err)
	}
	defer cc.Close()

	c := todopb.NewTodoServiceClient(cc)

	now := time.Now()
	todo := &todopb.Todo{
		UserId: "abc123",
		Title:  "buy",
		Detail: "buy a cup of coffee",
		Deadline: &timestamp.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}

	fmt.Println("createing the todo")

	res, err := c.CreateTodo(context.Background(), &todopb.CreateTodoRequest{
		Todo: todo,
	})
	if err != nil {
		log.Fatalf("failed to create todo: %v", err)

	}

	fmt.Printf("todo has been created: %v", res)

}
